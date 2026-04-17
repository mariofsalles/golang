package main

import (
	"custom-cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida da aplicação.")
	fmt.Println("..............................")
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	fmt.Println("..............................")
}
