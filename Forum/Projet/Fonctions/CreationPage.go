package Fonctions

import (
	dt "forum/DataBase"
	"html/template"
	"net/http"
)

var MessageERRInstance string

type DataCréation struct {
	Categorie     []Categorie
	MessageErreur string
}

func CreationPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		CreationGet(w, r)

	case "POST":
		CreationPost(w, r)
	}
}

func CreationGet(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./HTML/creation.html"))

	data := DataCréation{
		Categorie:     CreateTabCategorie(),
		MessageErreur: MessageERRInstance,
	}

	MessageERRInstance = ""

	tmpl.Execute(w, data)
}

func CreationPost(w http.ResponseWriter, r *http.Request) {
	//Récupération title et catégorie
	title := r.FormValue("Title")
	post := r.FormValue("Texte")
	r.ParseForm()
	categorie := r.FormValue("categorie_choice")

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}
	user := dt.GetUser(cookie.Value)

	if categorie == "" {
		MessageERRInstance = "Veuillez choisir une catégorie."
	}

	if MessageERRInstance == "" {
		dt.InsertPost(user.Id, title, post, categorie)
		dt.NbPostUp(user.Id)
		http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/Creation", http.StatusSeeOther)
}
