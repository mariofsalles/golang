package main

import "fmt"

func main() {
	// estrutura de chave e valor
	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}

	fmt.Println("Mapa inicial:", meuMap)
	fmt.Println("Valor da chave a:", meuMap["a"])
	fmt.Println("Valor da chave f:", meuMap["f"])

	// adicionando um novo valor
	meuMap["g"] = 7
	fmt.Println("Valor da chave g após adicionar:", meuMap["g"])
	fmt.Println("Mapa após adicionar a chave g:", meuMap)

	// deletando um valor
	delete(meuMap, "g")
	fmt.Println("Valor da chave g após deletar:", meuMap["g"])
	fmt.Println("Mapa final:", meuMap)


	fmt.Println("------Iterando sobre o mapa------")
	for chave, valor := range meuMap {
		fmt.Printf("Chave: %s, Valor: %v\n", chave, valor)
	}

	fmt.Println("------Mapa vazio------")
	meuMap = make(map[string]any)
	fmt.Println("Mapa após esvaziar:", meuMap)
	fmt.Println("Mapa de uma chave inexistente:", meuMap["a"])

	fmt.Println("------criar mapa------")
	meuMap["a"] = 1
	valor, ok := meuMap["a"].(string) // ok é um booleano que indica se a conversão foi bem-sucedida
	if ok {
		fmt.Println("Valor da chave a:", valor)
	} else {
		fmt.Println("Chave a não encontrada")
	}
}
