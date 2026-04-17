package main

import "fmt"

func fibonacci(x uint) uint {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {
	posicao := uint(10)
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
