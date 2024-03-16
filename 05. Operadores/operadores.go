package main

import "fmt"

func validaNumero(num int) string {
	var texto string
	if num > 10 {
		texto = "Maior que 10"
	} else {
		texto = "Menor que 10"
	}
	return texto
}

func main() {
	// Operadores Aritméticos
	// + - / * %
	// variaveis de tipos diferentes não podem ser operadas
	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mult := 10 * 5
	resto := 10 % 2
	println(soma, sub, div, mult, resto)
	// Operadores de Atribuição
	// = += -= *= /= %=
	var num int = 10
	num += 10
	fmt.Println(num)
	num -= 10
	fmt.Println(num)
	num *= 10
	fmt.Println(num)
	num /= 10
	fmt.Println(num)
	num %= 11
	fmt.Println(num)

	// Operadores Relacionais
	// == != >= <= > <
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)

	// Operadores Lógicos
	// && || !
	fmt.Println()
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println()
	// Operadores Unários
	// ++ --
	num1 := 10
	num1++ // num1++
	println(num1)

	num2 := 10
	num2--
	fmt.Println(num2)
	// Operadores Ternários
	// ?:
	fmt.Println()

	valida := validaNumero(11)
	fmt.Println(valida)
}
