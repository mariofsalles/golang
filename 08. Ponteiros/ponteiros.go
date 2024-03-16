package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// Ponteiro é uma referência de memória
	var var3 int
	var ponteiro *int

	fmt.Println()
	var3 = 100
	ponteiro = &var3
	fmt.Println(var3, &var3)
	fmt.Println()
	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro)

	var3 = 150
	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro)

}