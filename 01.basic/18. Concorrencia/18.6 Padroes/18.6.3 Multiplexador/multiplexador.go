package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo"), escrever("Programando em Go!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// recebe canais de entrada (mais de um canal) e retorna um canal de saída
func multiplexar(chanEntrada1, chanEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-chanEntrada1:
				canalSaida <- mensagem
			case mensagem := <-chanEntrada2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}
