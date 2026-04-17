package main

import (
	"fmt"
	"math"
)

// interface recebe assinatura de metodos
type forma interface {
	// aqui temos um metodo area que retorna um float64
	area() float64
	perimetro() float64
}

/*
Elementos geometricos que implementam a interface forma:
- Retangulo
- Circulo
*/
type retangulo struct {
	altura  float64
	largura float64
}

// a assinatura desse metodo deve ser igual ao que está na interface
// senão não será possivel implementar a interface
func (r retangulo) area() float64 {
	return r.altura * r.largura
}
func (r retangulo) perimetro() float64 {
	return 2 * (r.altura + r.largura)
}

type circulo struct {
	raio float64
}

// a assinatura desse metodo deve ser igual ao que está na interface
// senão não será possivel implementar a interface
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func escreverResultadosDaForma(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
	fmt.Printf("O perimetro da forma é %0.2f\n", f.perimetro())
	fmt.Println("-------------------------------------------------")
}

func main() {
	fmt.Println("Interfaces in Go")
	fmt.Println("-------------------------------------------------")

	r := retangulo{altura: 10, largura: 15}
	escreverResultadosDaForma(r)

	c := circulo{raio: 10}
	escreverResultadosDaForma(c)

}
