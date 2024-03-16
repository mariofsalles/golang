package main

import "fmt"

func diaDaSemana(num int) string {
	var diaDaSemana string
	switch {
	case num == 1:
		diaDaSemana = "Domingo"
	case num == 2:
		diaDaSemana = "Segunda-feira"
	case num == 3:
		diaDaSemana = "Terça-feira"
	case num == 4:
		diaDaSemana = "Quarta-feira"
	case num == 5:
		diaDaSemana = "Quinta-feira"
	case num == 6:
		diaDaSemana = "Sexta-feira"
	case num == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Dia da semana inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemana(2)
	fmt.Println(dia)
	dia = diaDaSemana(8)
	fmt.Println(dia)
}
