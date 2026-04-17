package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{10, 15}
		areaEsperada := float64(150)
		areaRecebida := retangulo.Area()
		if areaRecebida != areaEsperada {
			t.Fatalf("Area recebida %f, esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()
		if areaRecebida != areaEsperada {
			t.Fatalf("Area recebida %f, esperada %f", areaRecebida, areaEsperada)
		}
	})
}