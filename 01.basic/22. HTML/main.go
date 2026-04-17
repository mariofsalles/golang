package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		user := user{Name: "John Doe", Email: "joao@pedro.com"}

		templates.ExecuteTemplate(w, "home.html", user)
	})
	fmt.Println("Server is running at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
