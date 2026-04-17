package main

import "fmt"

func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}
	return
}

func escrever(texto string, num ...int) {  // o variatico deve ser o ultimo termo da funcao
	for _, v := range num {
		fmt.Println(texto, v)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 7, 8, 9, 10))
	escrever("Mario", 1, 2, 3, 4, 5)
}
