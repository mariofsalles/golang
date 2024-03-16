package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	c1 := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c1)
	fmt.Println("--------------")
	cachorroEmJSON, err := json.Marshal(c1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
	fmt.Println("------------------------------------------")
	c2 := map[string]string{
		"nome":  "Toby",
		"raca":  "Poodle",
		"idade": "5",
	}
	cachorroEmJSON, err = json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
}
