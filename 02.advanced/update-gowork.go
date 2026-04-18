package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	workDir := flag.String("dir", ".", "working directory for synchronization")
	flag.Parse()

	if err := os.Chdir(*workDir); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error changing to directory %s: %v\n", *workDir, err)
		os.Exit(1)
	}

	foundModules, err := scanModules()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error walking the project: %v\n", err)
		os.Exit(1)
	}

	_, err = os.Stat("go.work")
	goWorkExists := err == nil
	workModules, err := readGoWork("go.work")
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error reading go.work: %v\n", err)
		os.Exit(1)
	}

	toAdd, toRemove := diffModules(foundModules, workModules)

	if !goWorkExists && len(foundModules) == 0 {
		fmt.Println("⚠️  No go.mod files found.")
		return
	}

	if goWorkExists && len(toAdd) == 0 && len(toRemove) == 0 {
		fmt.Println("✅ go.work is already up to date.")
		return
	}

	if err := reconcileGoWork(goWorkExists, foundModules, toAdd, toRemove); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error updating go.work: %v\n", err)
		os.Exit(1)
	}

	if !goWorkExists {
		fmt.Println("🆕 Created go.work.")
	}
	for _, module := range toRemove {
		fmt.Printf("🗑️  Removed %s\n", module)
	}
	for _, module := range toAdd {
		fmt.Printf("➕ Added %s\n", module)
	}

	fmt.Println("✅ go.work updated.")
}

func scanModules() ([]string, error) {
	var modules []string
	seen := make(map[string]struct{})
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() != "go.mod" {
			return nil
		}

		dir := filepath.ToSlash(filepath.Clean(filepath.Dir(path)))
		if dir == "." {
			dir = "./"
		} else if !strings.HasPrefix(dir, "./") {
			dir = "./" + dir
		}
		if _, exists := seen[dir]; exists {
			return nil
		}
		seen[dir] = struct{}{}
		modules = append(modules, dir)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Strings(modules)
	return modules, nil
}

func diffModules(foundModules []string, workModules []string) ([]string, []string) {
	foundSet := make(map[string]struct{}, len(foundModules))
	for _, module := range foundModules {
		foundSet[module] = struct{}{}
	}

	workSet := make(map[string]struct{}, len(workModules))
	for _, module := range workModules {
		workSet[module] = struct{}{}
	}

	var toAdd []string
	for _, module := range foundModules {
		if _, ok := workSet[module]; !ok {
			toAdd = append(toAdd, module)
		}
	}

	var toRemove []string
	for _, module := range workModules {
		if _, ok := foundSet[module]; !ok {
			toRemove = append(toRemove, module)
		}
	}

	sort.Strings(toAdd)
	sort.Strings(toRemove)
	return toAdd, toRemove
}

func readGoWork(path string) ([]string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var workModules []string
	inUseBlock := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		switch {
		case trimmed == "use (":
			inUseBlock = true
			continue
		case inUseBlock && trimmed == ")":
			inUseBlock = false
			continue
		case inUseBlock:
			if strings.HasPrefix(trimmed, "./") {
				workModules = append(workModules, trimmed)
			}
			continue
		case strings.HasPrefix(trimmed, "use "):
			modulePath := strings.TrimSpace(strings.TrimPrefix(trimmed, "use "))
			if strings.HasPrefix(modulePath, "./") {
				workModules = append(workModules, modulePath)
			}
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return workModules, nil
}

func reconcileGoWork(goWorkExists bool, foundModules []string, toAdd []string, toRemove []string) error {
	goWorkPath, err := filepath.Abs("go.work")
	if err != nil {
		return err
	}

	if !goWorkExists {
		if len(foundModules) == 0 {
			return nil
		}
		if err := runGoCommand("off", "work", "init", foundModules[0]); err != nil {
			return err
		}
		if len(foundModules) == 1 {
			return nil
		}
		return runGoCommand(goWorkPath, append([]string{"work", "use"}, foundModules[1:]...)...)
	}

	if len(toRemove) > 0 {
		args := []string{"work", "edit"}
		for _, module := range toRemove {
			args = append(args, "-dropuse="+module)
		}
		if err := runGoCommand(goWorkPath, args...); err != nil {
			return err
		}
	}

	if len(toAdd) == 0 {
		return nil
	}

	return runGoCommand(goWorkPath, append([]string{"work", "use"}, foundModules...)...)
}

func runGoCommand(gowork string, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Env = append(os.Environ(), "GOWORK="+gowork)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
