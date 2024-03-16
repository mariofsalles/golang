package main

import "fmt"

var num int

func main() {
	fmt.Println("Funcao main")
	fmt.Println(num)
}

// pode ser utilizada para fazer um setup inicial
func init() {
	num = 20
	fmt.Println("Funcao init")
}
