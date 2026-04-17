package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 400)
			ch1 <- "1º"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			ch2 <- "2º"
		}
	}()

	for {
		select {
		case mensagem1 := <-ch1:
			fmt.Printf("mensagem1: %v\n", mensagem1)
		case mensagem2 := <-ch2:
			fmt.Printf("mensagem2: %v\n", mensagem2)
		}
	}
}
