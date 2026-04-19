package main

import "fmt"

func main() {
	// map -> [1,2,3,4,5] --> [2,4,6,8,10]
	// filter -> [1,2,3,4,5] --> [2,4]
	// reduce -> [1,2,3,4,5] --> 15
	var lista mySlice = []int{1, 2, 3, 4, 5}
	fmt.Println("Lista original:", lista)
	fmt.Println("--- Aplicando map, filter e reduce ---")

	resultadoFilter := lista.Filter(func(x int) bool {
		return x%2 == 0
	})
	
	resultadoMap := lista.Map(func(x int) int {
		return x * 2
	})

	resultadoReduce := lista.Reduce(func(acc, x int) int {
		return acc * x
	}, 1)
	
	fmt.Println("Resultado do filter:", resultadoFilter)
	fmt.Println("Resultado do map:", resultadoMap)
	fmt.Println("Resultado do reduce:", resultadoReduce)

	fmt.Println("--- Encadeamento de operações ---")
	resultadoEncadeado := lista.
		Filter(func(x int) bool { return x%2 == 1 }). // [1,3,5]
		Map(func(x int) int { return x * 3 }).        // [3,9,15]
		Map(func(x int) int { return x + 1 }).        // [4,10,16]
		Filter(func(x int) bool { return x > 5 }).    // [10,16]
		Reduce(func(acc, x int) int { return acc + x }, 0) // 26
	fmt.Println("Resultado do encadeamento:", resultadoEncadeado)
}

type mySlice []int

func (s mySlice) Map(f func(int) int) mySlice {
	var result mySlice
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

func (s mySlice) Filter(condition func(int) bool) mySlice {
	var result mySlice
	for _, value := range s {
		if condition(value) {
			result = append(result, value)
		}
	}
	return result
}

func (s mySlice) Reduce(f func(int, int) int, initial int) int {
	result := initial
	for _, value := range s {
		result = f(result, value)
	}
	return result
}
