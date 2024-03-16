package main

import (
	"fmt"
	"teste-avancado/formas"
)

func main() {
	r := formas.Retangulo{Altura: 10, Largura: 15}
	fmt.Printf("Retangulo:%v\n", r)
	formas.EscreverArea(r)

	c := formas.Circulo{Raio: 10}
	fmt.Printf("Circulo:%v\n", c)
	formas.EscreverArea(c)
}
