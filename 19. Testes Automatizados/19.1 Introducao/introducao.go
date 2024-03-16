package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoDeEndereco1 := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoDeEndereco1)
	tipoDeEndereco2 := enderecos.TipoDeEndereco("Praça das Rosas")
	fmt.Println(tipoDeEndereco2)
}
