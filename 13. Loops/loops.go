package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 200)
	}

	fmt.Println()

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Millisecond * 200)
	}

	fmt.Println()

	nomes := [3]string{"Mario", "Miguel", "Ceci"}

	for k, v := range nomes {
		nome := fmt.Sprintf(`posição[%v]: %s`, k, v)
		fmt.Println(nome)
		time.Sleep(time.Millisecond * 200)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
		time.Sleep(time.Millisecond * 200)
	}

	usuarios := map[string]string{
		"nome":      "Mario",
		"sobrenome": "Junior",
	}

	for chave, valor := range usuarios {
		fmt.Println(chave, valor)
		time.Sleep(time.Millisecond * 200)
	}

}
