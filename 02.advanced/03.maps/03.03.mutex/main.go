package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(100)
	for i := range 100 {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	
	for k,v := range m {
		fmt.Printf("Chave: %d, Valor: %d\n", k, v)
	}
}
