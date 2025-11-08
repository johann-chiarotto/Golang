package Fonctions

import (
	"fmt"
	dt "forum/DataBase"
	"html/template"
	"net/http"
)

var MessageErrInscription string

type dataInscription struct {
	MessageErreurInscription string
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		RegisterGet(w, r)

	case "POST":
		RegisterPost(w, r)
	}
}

func RegisterGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/register.html"))

	data := dataInscription{
		MessageErreurInscription: MessageErrInscription,
	}

	MessageErrInscription = ""

	tmpl.Execute(w, data)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // Gère s'il y a des erreurs
		fmt.Printf("Error ParseForm() : %v\n", err)
		http.Error(w, "Error while analyzing the form", http.StatusBadRequest)
		return
	}

	//Début de création du compte
	pseudoInscription := r.Form.Get("inputPseudoInscription")
	mdp1 := r.Form.Get("inputMdpInscription1")
	mdp2 := r.Form.Get("inputMdpInscription2")

	if mdp1 != mdp2 { // Vérification que le mot de passe soit le meme que celui de vérification
		MessageErrInscription = "Error both password need to be the same."
		http.Redirect(w, r, "/Register", http.StatusSeeOther)
		return
	}

	if len(pseudoInscription) > 10 { // Vérification que le pseudo ne soit pas trop grand
		MessageErrInscription = "Error the username has more than 10 characters"
		http.Redirect(w, r, "/Register", http.StatusSeeOther)
		return
	}

	tabUsersInscription := dt.GetUsers()

	for _, user := range tabUsersInscription {
		if user.Username == pseudoInscription {
			MessageErrInscription = "This username already exist, Use another one"
			http.Redirect(w, r, "/Register", http.StatusSeeOther)
			return
		}
	}

	dt.InsertUser(pseudoInscription, mdp1, 0)

	http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
}
