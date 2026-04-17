package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")
	
	var u usuario
	fmt.Println(u)

	end1 := endereco{"Rua dos Bobos", 0}

	u.nome = "Joca"
	u.idade = 20
	u.endereco = end1
	
	fmt.Println(u)


	u1 := usuario{
		nome:  "Pedro",
		idade: 20,
		endereco: end1,
	}

	fmt.Println(u1)
}
