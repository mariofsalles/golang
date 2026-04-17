package main

import (
	"errors"
	"fmt"
)

func main() {
	var n1 int8 = 8 // alias byte
	var n2 int16 = 16
	var n3 int32 = 32 // alias rune
	var n4 int64 = 64
	fmt.Println(n1, n2, n3, n4)

	var nreal1 float32 = 32.32
	var nreal2 float64 = 64.64
	fmt.Println(nreal1, nreal2)

	var str string = "Texto"
	fmt.Println(str)

	char := 'A' // tras o valor da tabela ASCII
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var numero int32
	fmt.Println(numero)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	var booleano1 bool
	fmt.Println(booleano1)

	var real float32
	fmt.Println(real)

}
