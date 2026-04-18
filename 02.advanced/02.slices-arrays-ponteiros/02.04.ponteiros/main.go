package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a

	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor de b (endereço de a):", b)
	fmt.Println("Valor apontado por b (valor de a):", *b)

	p := NovaPessoa("Mario", 25)
	fmt.Println("Pessoa:", p)

	
	fmt.Println("Telefone:", p.Telefone())

	fmt.Println("---- Definindo o telefone ----")	
	p.SetTelefone("1234-5678")
	p.SetIdade(26)
	fmt.Println("Pessoa e endereço memoria telefone:", p)
	fmt.Println("Telefone:", p.Telefone())
}

type Pessoa struct {
	Nome string
	Idade int
	telefone *string
}

func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome: nome,
		Idade: idade,
	}
}

func (p Pessoa) Telefone() string {
	if p.telefone == nil {
		return "Undefined"
	}
	return *p.telefone
}

func (p *Pessoa) SetTelefone(telefone string) {
	p.telefone = &telefone
}

func (p *Pessoa) SetIdade(idade int) {
	p.Idade = idade
}