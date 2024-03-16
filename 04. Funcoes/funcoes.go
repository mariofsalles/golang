package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	soma, sub := calculosMatematicos(10, 20)
	fmt.Println(soma, sub)

	var f = func (txt string) string {
		return txt
	}

	fmt.Println(f("Função f"))

}
