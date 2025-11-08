package Fonctions

import (
	"fmt"
	ch "forum/Chiffrement"
	dt "forum/DataBase"
	"html/template"
	"log"
	"net/http"
)

var ErreurInstance string

type DataProfil struct {
	Pdp          string
	Username     string
	Password     string
	DateCreation string
	NumbPosts    int
	NumbLikes    int
	NumbComments int
	Admin        bool
	Erreur       string
}

func changeDate(s string) string {
	var jour string
	var mois string
	var annee string

	for i := 0; i < len(s); i++ {
		if i < 4 {
			annee += string(s[i])
		}
		if i > 4 && i < 7 {
			mois += string(s[i])
		}
		if i > 7 && i < 10 {
			jour += string(s[i])
		}
	}

	return jour + " / " + mois + " / " + annee
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "user_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Expire immédiatement
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
}

func ProfilPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ProfilGet(w, r)

	case "POST":
		ProfilPost(w, r)
	}
}

func ProfilGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/profil.html"))

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}
	user := dt.GetUser(cookie.Value)

	IsAdmin := false
	if user.Admin == 1 {
		IsAdmin = true
	}

	data := DataProfil{
		Pdp:          user.Photo,
		Username:     user.Username,
		Password:     user.Password,
		DateCreation: changeDate(user.Created_at),
		NumbPosts:    user.NbPosts,
		NumbLikes:    user.NbLikes,
		NumbComments: user.NbComments,
		Admin:        IsAdmin,
		Erreur:       ErreurInstance,
	}

	ErreurInstance = ""

	tmpl.Execute(w, data)
}

func ProfilPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // Gère s'il y a des erreurs
		fmt.Printf("Erreur ParseForm() : %v\n", err)
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}
	user := dt.GetUser(cookie.Value)

	newUrl := r.Form.Get("url_pdp")

	if newUrl != "" {
		dt.ChangePdp(newUrl, user.Id)
	}

	// Change mdp
	mdpBase := r.Form.Get("password1")
	mdp1 := r.Form.Get("new_password1")
	mdp2 := r.Form.Get("new_password2")

	if mdpBase != "" && mdp1 != "" && mdp2 != "" {
		if mdp1 != mdp2 { // Vérification que le mot de passe soit le meme que celui de vérification
			ErreurInstance = "erreur les deux mots de passe ne sont pas les même."
			http.Redirect(w, r, "/Profil", http.StatusSeeOther)
			return
		}

		if ch.CheckPasswordHash(mdpBase, user.Password) {
			mdp, err := ch.HashPassword(mdp1)
			if err != nil {
				log.Fatal(err)
			}
			dt.ChangePassword(mdp, user.Id)
			ErreurInstance = "Mot de passe mit a jour"
		} else {
			ErreurInstance = "Mauvais mot de passe."
		}
	}
	//Fin change mdp

	http.Redirect(w, r, "/Profil", http.StatusSeeOther)
}
