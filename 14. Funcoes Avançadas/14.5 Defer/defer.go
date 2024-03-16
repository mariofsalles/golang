package main

import "fmt"

func funcao1() {
	println("Executando a função 1")
}

func funcao2() {
	println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer println("Média calculada. Resultado será retornado") // será executado antes do retorno
	println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	funcao1() // será adiada para o final da execução da função main
	funcao2()
	fmt.Println()
	fmt.Println(alunoEstaAprovado(7, 8))
	fmt.Println()
	fmt.Println(alunoEstaAprovado(5, 5))
	fmt.Println()
}