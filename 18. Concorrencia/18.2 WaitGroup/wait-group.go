package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(4) // rodam 4 goroutines ao mesmo tempo

	// primeira goroutine
	go func() {
		escrever("Olá Mundo")
		wg.Done() // decrementa o contador em wg.Add
	}()

	// segunda goroutine
	go func() {
		escrever("Programando em Go")
		wg.Done() // decrementa o contador em wg.Add
	}()

	// terceira goroutine
	go func() {
		escrever("Goroutine 3")
		wg.Done() // decrementa o contador em wg.Add
	}()

	// quarta goroutine
	go func() {
		escrever("Goroutine 4")
		wg.Done() // decrementa o contador em wg.Add
	}()

	wg.Wait() // espera o contador chegar a 0, ou seja espera as duas goroutines terminarem
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Millisecond * 500)
	}
}
