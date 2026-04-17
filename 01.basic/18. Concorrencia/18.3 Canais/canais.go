package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Canais")
	fmt.Println("-----------------")
	canal := make(chan string)
	go escrever("Olá Mundo", canal) // tem um script que envia um texto para o canal
	fmt.Println("Depois da função escrever começar a ser executada")
	// for {
	// 	mensagem, aberto := <-canal // recebe o texto enviado pela goroutine escrever
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // envia o texto para o canal
		time.Sleep(time.Millisecond * 500)
	}
	close(canal) // fecha o canal para que não ocorra um deadlock
}
