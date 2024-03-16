package main

import "fmt"

func main() {
	var var1 string = "Variavel 1"
	var2 := "Variavel 2"
	fmt.Println(var1)
	fmt.Println(var2)

	var (
		var3 string = "Variavel 3"
		var4 string = "Variavel 4"
	)
	fmt.Println(var3, var4)

	var5, var6 := "Variavel 5", "Variavel 6"
	fmt.Println(var5, var6)

}
