package main

import "fmt"

func main() {
	var (
		array = [5]int{1, 2, 3, 4, 5}
		slice = []int{1, 2, 3, 4, 5}
	)

	fmt.Printf("Array: len=%d cap=%d\n", len(array), cap(array))
	fmt.Printf("Slice: len=%d cap=%d\n", len(slice), cap(slice))

	fmt.Println("---- Modificando o slice ----")
	slice = append(slice, 6)
	fmt.Printf("Slice after append: len=%d cap=%d\n", len(slice), cap(slice))
	// array = append(array, 6) // isso nao compila, pois array tem tamanho fixo

	fmt.Println("---- Imprimindo array e slice ----")
	printSlice(slice)
	// printArray(array) // isso nao compila, pois array e slice tem tipos diferentes
	fmt.Printf("Array: %v\n", array)

	fmt.Println("---- Modificando o slice novamente ----")
	slice = append(slice, array[:]...)
	fmt.Printf("Slice after append: len=%d cap=%d\n", len(slice), cap(slice))
	fmt.Printf("Slice after appending array: %v\n", slice)
}

func printSlice(s []int) {
	for i, v := range s {
		fmt.Printf("Slice[%d]: %d\n", i, v)
	}
}
