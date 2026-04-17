// concorrencia é a capacidade de executar várias tarefas ao mesmo tempo
// paralelismo é a capacidade de executar várias tarefas ao mesmo tempo em diferentes processadores
package main

import (
	"fmt"
	"time"
)

func main() {
	// gorroutines executa o codigo e não espera que ele termine para executar o proximo
	go escrever("Olá Mundo")
	go escrever("Programando em Go")
	escrever("Proximo")
}
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
