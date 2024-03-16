package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	var array1 [3]int = [3]int{1, 2, 3}
	fmt.Println(array1)

	array2 := [...]int{1, 2, 3, 4}
	// array2[4] = 5, não funciona isso
	fmt.Println(array2)

	var slice1 []int
	slice1 = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice1)

	fmt.Println()

	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(slice1))

	// Arrays Internos
	slice3 := make([]int, 10, 11)
	fmt.Println("--------------------")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 1)
	fmt.Println("--------------------")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("--------------------")
	slice4 := make([]int, 3)
	array4 := []int{1, 2, 3}
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, array4...)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
