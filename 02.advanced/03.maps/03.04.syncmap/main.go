package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map
		wg sync.WaitGroup
	)

	wg.Add(100)
	for i := range 100 {
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}

	wg.Wait()
	m.Range(func(key, value any) bool {
		fmt.Println("Chave:", key, "Valor:", value)
		return true
	})

	fmt.Println("------Carregando valor específico------")
	valor, ok := m.Load(50)
	if ok {
		fmt.Println("Valor carregado:", valor)
	}
}