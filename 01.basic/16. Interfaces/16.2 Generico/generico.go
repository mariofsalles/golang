package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(1.5)
	generica(true)
	generica([]int{1, 2, 3})
	generica(map[int]string{1: "a", 2: "b"})
}
