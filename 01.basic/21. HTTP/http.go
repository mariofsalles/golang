package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}
func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página de usuários"))
}

func main() {
	// http protocolo de comunicação - baseado em requisição e resposta
	// requisição é feita por um cliente e a resposta é dada por um servidor
	// a requisição é uma mensagem
	// rotas são os caminhos que a requisição pode seguir com metodos que podem ser usados
	// a rota posui URI identificando o recurso
	// a rota possui um metodo que identifica a ação que será realizada: GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)
	
	log.Fatal(http.ListenAndServe(":3000", nil))
}
