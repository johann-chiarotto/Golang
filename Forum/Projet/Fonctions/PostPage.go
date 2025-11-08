package Fonctions

import (
	db "forum/DataBase"
	"html/template"
	"net/http"
	"strconv"
)

type DataPost struct {
	Admin          bool
	Post           db.Post
	Comments       []db.Comment
	Pdp_User       string
	UserLikeStatus int
}

func PostPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		PostGet(w, r)

	case "POST":
		PostPost(w, r)
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/post.html"))

	id := r.URL.Query().Get("id")

	post := db.GetPost(id)

	Pdp_User_instance := db.GetPdp(post.UserID)

	cookie, err := r.Cookie("user_id")
	user := db.GetUser(cookie.Value)

	IsAdmin := false

	if user.Admin == 1 {
		IsAdmin = true
	}

	var userLikeStatus int
	if err == nil {
		user := db.GetUser(cookie.Value)
		userLikeStatus = db.GetUserLikeStatus(post.ID, user.Id)
	}

	data := DataPost{
		Post:           post,
		Comments:       db.GetCommentsPost(id),
		Pdp_User:       Pdp_User_instance,
		UserLikeStatus: userLikeStatus,
		Admin:          IsAdmin,
	}

	tmpl.Execute(w, data)
}

func PostPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Récupère l'ID du Post

	if err := r.ParseForm(); err != nil { // Formulaire qui vérifie qu'il n'y ai pas d'errreures
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}
	post := db.GetPost(id) // Récupère le bon post

	cookie, err := r.Cookie("user_id") // Récupère le cookie de l'user connecter
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}
	UserConnected := db.GetUser(cookie.Value)

	// Ajout d'un commentaire
	CommentSupp := r.Form.Get("reponse")
	if CommentSupp != "" {
		db.InsertComment(UserConnected.Id, id, CommentSupp)
		db.NbCommentsUp(UserConnected.Id)
	}

	// Suppression d'un commentaire (admin)
	commentId := r.FormValue("commentId")
	if commentId != "" {
		db.DeleteComment(commentId)
		http.Redirect(w, r, "Post?id="+strconv.Itoa(post.ID), http.StatusSeeOther)
		return
	}

	// Suppression d'un post (admin)
	postId := r.FormValue("postId")
	if postId != "" {
		db.DeletePost(postId)
		db.DeleteComment(postId)
		db.DeleteLikes(postId)
		http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "Post?id="+strconv.Itoa(post.ID), http.StatusSeeOther)
}
