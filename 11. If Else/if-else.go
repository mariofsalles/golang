package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if v := numero; v > 0 { // v é uma variavel local limitada ao escopo deste if
		fmt.Println("Numero é maior que zero")
	} else if v < -10 {
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}
