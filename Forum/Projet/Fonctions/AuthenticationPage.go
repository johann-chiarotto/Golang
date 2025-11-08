package Fonctions

import (
	"fmt"
	ch "forum/Chiffrement"
	dt "forum/DataBase"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

var MessageErrConnexion string

type DataAuthentication struct {
	MessageErreurConnexion string
}

func AuthenticationPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		AuthenticationGet(w, r)

	case "POST":
		AuthenticationPost(w, r)
	}
}

func AuthenticationGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/authentication.html"))

	data := DataAuthentication{
		MessageErreurConnexion: MessageErrConnexion,
	}

	MessageErrConnexion = ""

	tmpl.Execute(w, data)
}

func AuthenticationPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Printf("Erreur ParseForm() : %v\n", err)
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	pseudoConnexion := r.Form.Get("inputPseudoConnexion")
	mdp := r.Form.Get("inputMdpConnexion")

	if pseudoConnexion != "" && mdp != "" {
		tabUsersConnexion := dt.GetUsers()
		for _, user := range tabUsersConnexion {
			if user.Username == pseudoConnexion && ch.CheckPasswordHash(mdp, user.Password) {
				cookie := http.Cookie{
					Name:     "user_id",
					Value:    strconv.Itoa(user.Id),
					Expires:  time.Now().Add(24 * time.Hour),
					HttpOnly: true,
				}
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
				return
			}
		}
		MessageErrConnexion = "Identifiants incorrects"
	}
	http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
}
