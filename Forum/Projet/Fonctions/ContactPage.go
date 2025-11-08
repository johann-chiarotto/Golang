package Fonctions

import (
	"html/template"
	"net/http"
)

type DataContact struct {
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/contact.html"))

	data := DataContact{}

	tmpl.Execute(w, data)
}
