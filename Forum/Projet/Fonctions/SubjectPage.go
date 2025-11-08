package Fonctions

import (
	db "forum/DataBase"
	"html/template"
	"net/http"
)

type DataSubject struct {
	PostsTab      []db.Post
	Categorie     []Categorie
	CategorieName string
}

func SubjectPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/subject.html"))

	categorie := r.URL.Query().Get("categorie")

	cookie, err := r.Cookie("user_id")
	var userID int
	if err == nil {
		user := db.GetUser(cookie.Value)
		userID = user.Id
	}

	posts := db.GetPostsCategorie(categorie)
	for i := range posts {
		posts[i].UserLikeStatus = db.GetUserLikeStatus(posts[i].ID, userID)
	}

	data := DataSubject{
		PostsTab:      posts,
		Categorie:     CreateTabCategorie(),
		CategorieName: categorie,
	}

	tmpl.Execute(w, data)
}
