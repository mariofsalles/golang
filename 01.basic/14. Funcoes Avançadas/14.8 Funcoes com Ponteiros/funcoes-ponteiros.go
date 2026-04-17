package main

import "fmt"

// passa copia de um valor para a função
func inverterSinal(num int) int {
	return -1 * num
}

// essa função altera diretamente o valor da variavel dessa forma não é necessário um retorno
// passamos uma referencia para a função
func inverterSinalComPonteiro(num *int) {
	*num = -1 * (*num) 
}

func main() {
	num := 10
	numero_a_inverter := inverterSinal(num)
	fmt.Println(numero_a_inverter)

	// a variavel é a mesma, mas o valor é alterado devido a função inverterSinalComPonteiro
	novo_num := 20

	fmt.Println(novo_num)

	fmt.Println("-----------------")
	inverterSinalComPonteiro(&novo_num)
	fmt.Println(novo_num)

	fmt.Println("-----------------")
	inverterSinalComPonteiro(&novo_num)
	fmt.Println(novo_num)
}
