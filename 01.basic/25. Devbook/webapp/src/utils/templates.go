package utils

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

// LoadTemplates loads all templates in the views folder
func LoadTemplates(patterns []string) {
	var err error
	templates, err = template.New("").ParseGlob(patterns[0])
	if err != nil {
		log.Fatalf("Error loading templates: %v", err)
	}

	for _, pattern := range patterns[1:] {
		_, err := templates.ParseGlob(pattern)
		if err != nil {
			log.Fatalf("Error parsing additional templates: %v", err)
		}
	}
}

// ExecuteTempate renders a html page with the given data
func ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) {
	templates.ExecuteTemplate(w, name, data)
}
