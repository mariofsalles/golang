package main

import (
	"encoding/json"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`
	var c1 cachorro
	if err := json.Unmarshal([]byte(cachorroEmJSON), &c1); err != nil {
		log.Fatal(err)
	}
	log.Println(c1)

	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle","idade":"5"}`
	c2 := make(map[string]string)
	if err := json.Unmarshal([]byte(cachorro2EmJSON), &c2); err != nil {
		log.Fatal(err)
	}
	log.Println(c2)

}
