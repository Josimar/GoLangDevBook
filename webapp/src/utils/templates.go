package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates -> pega os html e carregam na variável template
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate -> renderiza o template html
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	err := templates.ExecuteTemplate(w, template, dados)
	if err != nil {
		return
	}
}
