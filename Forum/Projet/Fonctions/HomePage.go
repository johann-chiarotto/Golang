package Fonctions

import (
	dt "forum/DataBase"
	"html/template"
	"net/http"
)

var tabPostsInstance []dt.Post

type Categorie struct {
	Name  string
	Image string
}

type Data struct {
	IsConnect bool
	Username  string
	PostsTab  []dt.Post
	Categorie []Categorie
	Photo_pdp string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		HomePageGet(w, r)

	case "POST":
		HomePagePost(w, r)
	}
}

func HomePageGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./HTML/homepage.html"))

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/Authentication", http.StatusSeeOther)
		return
	}
	user := dt.GetUser(cookie.Value)

	posts := dt.GetPosts()
	for i := range posts {
		posts[i].UserLikeStatus = dt.GetUserLikeStatus(posts[i].ID, user.Id)
	}

	data := Data{
		IsConnect: true,
		Username:  user.Username,
		PostsTab:  ReverseSlice(posts),
		Categorie: CreateTabCategorie(),
		Photo_pdp: user.Photo,
	}
	tmpl.Execute(w, data)
}

func HomePagePost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
}

func CreateTabCategorie() []Categorie {
	var NameCat = []string{"Animals", "Cuisine", "Video Games", "Famous Historical Figures", "Music", "Sports", "Vehicles", "Other / Miscellaneous"}
	var ImageCat = []string{"https://thumbs.dreamstime.com/b/cat-et-chien-dormant-le-sommeil-des-chiots-chats-156181857.jpg", "https://images.radio-canada.ca/q_auto,w_960/v1/ici-premiere/16x9/bien-entendu-cuisiner-instinct-intuition-salade-cuisine-plaisir.jpg",
		"https://f.hellowork.com/blogdumoderateur/2024/04/chiffres-industrie-jeu-video.jpg", "https://i.ytimg.com/vi/d1DFRFjfaK0/maxresdefault.jpg", "https://www.profession-spectacle.com/wp-content/uploads/2018/12/Saxophone-jazz-band-1280x640.jpg",
		"https://img.freepik.com/photos-gratuite/outils-sport_53876-138077.jpg", "https://www.trajetalacarte.com/blog/wp-content/uploads/2024/05/Les-10-voitures-les-plus-populaires-en-Inde-1024x683.jpg",
		"https://media.istockphoto.com/id/1132309800/fr/photo/point-dinterrogation-apprentissage-recherche-pens%C3%A9e-homme-demandant-homme-avec-orange-3d.jpg?s=612x612&w=0&k=20&c=Ng1Iq3N0gTFX-f0VUuccJzqp-Ed9sskiU3DWkZ3hUqc="}

	var CategoriesFinal []Categorie

	for i := 0; i < len(NameCat); i++ {
		var CategorieInstance Categorie
		CategorieInstance.Name = NameCat[i]
		CategorieInstance.Image = ImageCat[i]

		CategoriesFinal = append(CategoriesFinal, CategorieInstance)
	}

	return CategoriesFinal
}

func SetupRoutes() {
	http.HandleFunc("/logout", LogoutHandler)
}

func RedirectBasePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
}

func ReverseSlice(tab []dt.Post) []dt.Post {
	var TabInverse []dt.Post

	for i := len(tab) - 1; i > -1; i-- {
		TabInverse = append(TabInverse, tab[i])
	}

	return TabInverse
}
