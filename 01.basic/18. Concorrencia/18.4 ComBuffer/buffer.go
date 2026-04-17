package main

import "fmt"

func main() {
	canal := make(chan string, 3)

	canal <- "Olá Mundo1"
	canal <- "Olá Mundo2"
	canal <- "Olá Mundo3"

	mensagem1 := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Printf("mensagem1: %v\n", mensagem1)
	fmt.Printf("mensagem2: %v\n", mensagem2)
	fmt.Printf("mensagem3: %v\n", mensagem3)
}
