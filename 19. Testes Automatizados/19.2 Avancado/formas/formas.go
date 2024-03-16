package formas

import (
	"fmt"
	"math"
)

// Retangulo é uma struct que representa um retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Retangulo a assinatura desse metodo deve ser igual ao que está na interface
// senão não será possivel implementar a interface
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// Circulo é uma struct que representa um circulo
type Circulo struct {
	Raio float64
}

// Circulo a assinatura desse metodo deve ser igual ao que está na interface
// senão não será possivel implementar a interface
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

// Forma interface recebe assinatura de metodos
type Forma interface {
	// aqui temos um metodo area que retorna um float64
	Area() float64
}

// EscreverArea escreve a area da forma
func EscreverArea(f Forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.Area())
	fmt.Println("-------------------------------------------------")
}
