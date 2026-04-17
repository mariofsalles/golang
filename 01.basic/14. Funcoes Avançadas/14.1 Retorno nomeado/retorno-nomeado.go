package main

func operacoes(a, b int) (soma int, subtracao int) {
	soma = a + b
	subtracao = a - b
	return
}

func main() {
	soma, subtracao := operacoes(10, 20)
	println(soma, subtracao)
}
