package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = slice1 // isso nao cria uma copia, slice2 aponta para o mesmo array de slice1
	
	fmt.Println("Slice 1:", slice1) // Slice 1: [10 2 3 4 5]
	fmt.Println("Slice 2:", slice2) // Slice 2: [10 2 3 4 5]

	// inserir linha de separação
	fmt.Println("---- Modificar slice2, modifica o slice1 ----")
	slice2[0] = 10 // isso modifica o array compartilhado, afetando ambos os slices
	fmt.Println("Slice 1:", slice1) // Slice 1: [10 2 3 4 5]
	fmt.Println("Slice 2:", slice2) // Slice 2: [10 2 3 4 5]

	fmt.Println("---- Criando uma cópia utilizando append ----")
	slice2 = append(slice2, len(slice1)) // isso cria um novo array para slice2, pois a capacidade é excedida
	slice2[0] = 20 // isso modifica apenas slice2, pois é uma cópia independente
	fmt.Println("Slice 1:", slice1) // Slice 1: [10 2 3 4 5]
	fmt.Println("Slice 2:", slice2) // Slice 2: [20 2 3 4 5]

	fmt.Println("---- Criando uma cópia utilizando copy ----")
	slice3 := make([]int, len(slice1)) // criar um novo slice com o mesmo comprimento
	copy(slice3, slice1) // copiar os elementos de slice1 para slice3
	fmt.Println("Slice 1:", slice1) // Slice 1: [10 2 3 4 5]
	fmt.Println("Slice 3:", slice3) // Slice 3: [30 2 3 4 5]

	// ncluir uma linha de separação
	fmt.Println("---- Modificando slice3 ----")
	slice3[0] = 30 // isso modifica apenas slice3, pois é uma cópia independente
	fmt.Println("Slice 1:", slice1) // Slice 1: [10 2 3 4 5]
	fmt.Println("Slice 3:", slice3) // Slice 3: [30 2 3 4 5]
}
