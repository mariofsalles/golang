package main

import "fmt"

type Pessoa struct {
	Nome  *string
	Idade int
}

func main() {
	pessoaOriginal := Pessoa{
		Nome:  stringPtr("João"),
		Idade: 30,
	}

	copiaSuperficial := pessoaOriginal

	fmt.Println("--- Cópia superficial ---")
	imprimirPessoa("Pessoa original:", pessoaOriginal)
	imprimirPessoa("Cópia superficial:", copiaSuperficial)

	*pessoaOriginal.Nome = "Maria"
	fmt.Println("--- Alterando o nome pela struct original ---")
	imprimirPessoa("Pessoa original:", pessoaOriginal)
	imprimirPessoa("Cópia superficial:", copiaSuperficial)

	fmt.Println("--- Cópia Profunda ---")
	copiaProfunda := copiarPessoaProfundamente(pessoaOriginal)
	*copiaProfunda.Nome = "Mario"
	imprimirPessoa("Pessoa original:", pessoaOriginal)
	imprimirPessoa("Cópia profunda:", copiaProfunda)

	// copia profunda de uma lista de pessoas
	fmt.Println("--- Cópia Profunda de uma lista de pessoas ---")
	pessoas := []Pessoa{
		{Nome: stringPtr("Alice"), Idade: 25},
		{Nome: stringPtr("Bob"), Idade: 28},
	}

	copiaPessoas := copiarPessoasProfundamente(pessoas)

	for i := range pessoas {
		imprimirPessoa(fmt.Sprintf("Pessoa original[%d]:", i), pessoas[i])
		imprimirPessoa(fmt.Sprintf("Cópia profunda[%d]:", i), copiaPessoas[i])
	}

	fmt.Println("--- Verificando endereços de memória ---")
	fmt.Println("Endereço da pessoa original:", pessoas)
	fmt.Println("Endereço da cópia profunda:", copiaPessoas)

}

func copiarPessoaProfundamente(origem Pessoa) Pessoa {
	copia := origem
	if origem.Nome != nil {
		copia.Nome = stringPtr(*origem.Nome)
	}
	return copia
}

func copiarPessoasProfundamente(origem []Pessoa) []Pessoa {
	copia := make([]Pessoa, len(origem))
	for i, pessoa := range origem {
		copia[i] = copiarPessoaProfundamente(pessoa)
	}
	return copia
}

func imprimirPessoa(rotulo string, pessoa Pessoa) {
	nome := "<nil>"
	if pessoa.Nome != nil {
		nome = *pessoa.Nome
	}
	fmt.Println(rotulo, nome, pessoa.Idade)
}

func stringPtr(s string) *string {
	return &s
}

