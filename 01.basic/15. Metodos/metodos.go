package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

// temos a funcao salvar que pertence a struct pessoa, p é a variavel
// que pode referenciar outros campos da struct
func (p pessoa) salvar() {
	fmt.Printf("Salvando o usuario %s no banco de dados\n", p.nome)
}

func (p pessoa) verificaIdade() string {
	var validacao string
	if p.idade >= 18 {
		validacao = fmt.Sprintf("%s é maior de idade e tem %d anos", p.nome, p.idade)
	} else {
		validacao = fmt.Sprintf("%s é menor de idade e tem %d anos", p.nome, p.idade)
	}
	return validacao
}

// se o metodo alterar o valor de algum campo da struct, é necessario
// passar o ponteiro da struct como parametro
func (p *pessoa) fazAniversario() {
	p.idade++
	fmt.Printf("%s fez aniversario e agora tem %d anos\n", p.nome, p.idade)
}

func main() {
	u1 := pessoa{"James", 32}
	u2 := pessoa{"Moneypenny", 17}

	u1.salvar()
	fmt.Println(u1.verificaIdade())
	fmt.Println()
	u2.salvar()
	fmt.Println(u2.verificaIdade())
	u2.fazAniversario()
	fmt.Println(u2.verificaIdade())
}
