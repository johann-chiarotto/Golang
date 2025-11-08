package Fonctions

import (
	dt "forum/DataBase"
	"html/template"
	"net/http"
)

type DataAdmin struct {
	Admin    bool
	UsersTab []dt.User
}

func GestionAdminPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GestionAdminGet(w, r)

	case "POST":
		GestionAdminPost(w, r)
	}
}

func GestionAdminGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/gestion_admin.html"))

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}
	user := dt.GetUser(cookie.Value) //Récupération de l'user connecté à l'aide du cookie

	IsAdmin := false
	var UsersTabInstance []dt.User

	if user.Admin == 1 {
		UsersTabInstance = dt.GetUsers()
		IsAdmin = true
	}

	data := DataAdmin{
		UsersTab: UsersTabInstance,
		Admin:    IsAdmin,
	}

	tmpl.Execute(w, data)
}

func GestionAdminPost(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userId") // Récupère l'id de la corbeille presser

	dt.DeleteUser(userID)

	http.Redirect(w, r, "/GestionAdmin", http.StatusSeeOther)
}
