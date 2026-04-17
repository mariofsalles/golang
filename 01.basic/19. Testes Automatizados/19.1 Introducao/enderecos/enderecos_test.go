// o arquivo deve possuir o sufixo _test no nome do arquivo
package enderecos_test

import (
	// o alias . é usado para o arquivo de teste na sua maioria das vezes
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Variavel que armazena os cenários de teste
var cenarios []cenarioDeTeste = []cenarioDeTeste{
	// Valid types to test
	{"Rua ABC", "Rua"},
	{"Avenida Paulista", "Avenida"},
	{"Estrada Qualquer", "Estrada"},
	{"Rodovia dos Imigrantes", "Rodovia"},
	// Invalid types to test
	{"Praça da Sé", "Tipo Inválido"},
	{"Viela dos Pássaros", "Tipo Inválido"},
	{"", "Tipo Inválido"},
}

// deve iniciar com "Test" concatenado com o nome da função que será testada
// e receber um parâmetro do tipo *testing.T
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	for _, cenario := range cenarios {
		if TipoDeEndereco(cenario.enderecoInserido) != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s.",
				cenario.retornoEsperado,
				TipoDeEndereco(cenario.enderecoInserido),
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Erro!")
	}
}

// testar a cobertura de testes: go test -cover
// testar a cobertura de testes e exibir o resultado no navegador: go test -coverprofile=resultado.out && go tool cover -html=resultado.out
// testar a cobertura de testes e exibir no terminal: go test -coverprofile=resultado.out && go tool cover -func=resultado.out
