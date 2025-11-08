package Fonctions

import (
	"fmt"
	db "forum/DataBase"
	"net/http"
	"strconv"
)

func SetDB() {
	db.CreateTableComments()
	db.CreateTablePosts()
	db.CreateTableLikes()
	db.CreateTableUsers()
	fmt.Println("")
}

func SetHttp() {
	fs := http.FileServer(http.Dir("./CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	http.HandleFunc("/", RedirectBasePage)
	http.HandleFunc("/HomePage", HomePage)
	http.HandleFunc("/Authentication", AuthenticationPage)
	http.HandleFunc("/Profil", ProfilPage)
	http.HandleFunc("/Subject", SubjectPage)
	http.HandleFunc("/Contact", ContactPage)
	http.HandleFunc("/GestionAdmin", GestionAdminPage)
	http.HandleFunc("/FAQ", FAQPage)
	http.HandleFunc("/Register", RegisterPage)
	http.HandleFunc("/Creation", CreationPage)
	http.HandleFunc("/Post", PostPage)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/like", HandleLike)
}

func HandleLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}
	user := db.GetUser(cookie.Value)

	action := r.FormValue("action")

	// Vérifier si l'utilisateur a déjà liké ce post
	var existingLike *db.Likes
	for _, like := range db.GetLikes() {
		if like.Post_id == postID && like.User_id == user.Id {
			existingLike = &like
			break
		}
	}

	if action == "like" {
		if existingLike != nil {
			if existingLike.Status == 1 {
				// L'utilisateur avait déjà liké -> supprimer le like
				db.DeleteLike(strconv.Itoa(existingLike.ID))
				db.NbLikes(postID, -1)
			} else {
				// L'utilisateur avait disliké -> supprimer le dislike et ajouter un like
				db.DeleteLike(strconv.Itoa(existingLike.ID))
				db.InsertLike(postID, user.Id)
				db.NbLikes(postID, 2)
			}
		} else {
			// L'utilisateur n'a jamais liké ce post
			db.InsertLike(postID, user.Id)
			db.NbLikes(postID, 1)
		}
	} else if action == "dislike" {
		if existingLike != nil {
			if existingLike.Status == -1 {
				// L'utilisateur avait déjà disliké -> supprimer le dislike
				db.DeleteLike(strconv.Itoa(existingLike.ID))
				db.NbLikes(postID, +1)
			} else {
				// L'utilisateur avait liké -> supprimer le like et ajouter un dislike
				db.DeleteLike(strconv.Itoa(existingLike.ID))
				db.InsertDislike(postID, user.Id)
				db.NbLikes(postID, -2)
			}
		} else {
			// L'utilisateur n'a jamais disliké ce post
			db.InsertDislike(postID, user.Id)
			db.NbLikes(postID, -1)
		}
	} else {
		http.Error(w, "Action invalide", http.StatusBadRequest)
		return
	}

	post := db.GetPost(strconv.Itoa(postID))
	fmt.Fprintf(w, "%d", post.Likes)
}
