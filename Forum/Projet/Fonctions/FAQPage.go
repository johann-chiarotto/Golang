package Fonctions

import (
	"html/template"
	"net/http"
)

type DataFaq struct {
}

func FAQPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./HTML/faq.html"))

	data := DataFaq{}

	tmpl.Execute(w, data)
}
