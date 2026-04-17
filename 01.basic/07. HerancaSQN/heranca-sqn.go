package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// nessa struct tem a herança só que não
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

type profissional struct {
	pessoa  pessoa
	empresa string
	funcao  string
}

func main() {
	fmt.Println("Herança SQN")

	p := pessoa{}	
	p.nome = "João"
	p.sobrenome = "Pedro"
	p.idade = 20
	p.altura = 178

	fmt.Println(p)

	e1 := estudante{}
	e1.nome = "João"
	e1.sobrenome = "Sousa"
	e1.idade = 20
	e1.altura = 178
	e1.curso = "Engenharia de Software"
	e1.faculdade = "UFSCar"
	fmt.Println(e1)

	p2 := profissional{}
	p2.pessoa.nome = "Maria"
	p2.pessoa.sobrenome = "Silva"
	p2.pessoa.idade = 24
	p2.pessoa.altura = 165

	p2.empresa = "Google"
	p2.funcao = "Engenheira de Software"
	
	fmt.Println(p2)

}
