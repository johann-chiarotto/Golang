package DataBase

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "modernc.org/sqlite"
)

// Structure qui reprend les caractéristiques de la Base de donnée
type Post struct {
	ID             int
	UserID         int
	Title          string
	Content        string
	Likes          int
	Categorie      string
	UserName       string
	UserLikeStatus int
}

// Fonction qui permet de retourner la connexion a la base de donnée posts
func connect_postsDB() *sql.DB {
	dbPosts, err := sql.Open("sqlite", "../DataBase/Stockage/Posts.db")
	if err != nil {
		panic(err)
	}

	// Vérifier la connexion
	if err := dbPosts.Ping(); err != nil {
		panic("Impossible de se connecter à la base de données dbPosts !")
	}

	return dbPosts
}

// Fonction qui créer la base de données posts si elle n'existe pas
// Si celle ci existe déjà alors elle reprend toutes les informations qui existes
func CreateTablePosts() {
	db := connect_postsDB()
	defer db.Close()

	// Structure de notre base de donnée
	query := `
	CREATE TABLE IF NOT EXISTS posts (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXTE NOT NULL,
		content TEXT NOT NULL,
		likes INTEGER DEFAULT 0,
		categorie TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table 'posts' créée avec succès.")
}

// Fonction qui renvoie un tableau de tous les posts qui sont dans la base de donnée posts
func GetPosts() []Post {
	db := connect_postsDB()

	rows, err := db.Query("SELECT id, user_id, title,content, likes, categorie FROM posts;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Final []Post
	for rows.Next() {
		var PostFinal Post

		if err := rows.Scan(&PostFinal.ID, &PostFinal.UserID, &PostFinal.Title, &PostFinal.Content, &PostFinal.Likes, &PostFinal.Categorie); err != nil {
			panic(err)
		}
		PostFinal.UserName = GetName(PostFinal.UserID)

		Final = append(Final, PostFinal)
	}
	return Final
}

// Fonction qui renvoie le post correspondant a l'id qui est pris en paramètre
func GetPost(id string) Post {
	db := connect_postsDB()

	query := `SELECT id, user_id, title, content, likes, categorie FROM posts WHERE id = ?;`

	var post Post
	err := db.QueryRow(query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Likes, &post.Categorie)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du post :", err)
		return Post{}
	}
	post.UserName = GetName(post.UserID)

	return post
}

// Fonction qui renvoie un tableau de tous les posts qui correspondent a une catégorie donner en paramètre
func GetPostsCategorie(categorie string) []Post {
	db := connect_postsDB()

	rows, err := db.Query("SELECT id, user_id, title, content, likes, categorie FROM posts WHERE categorie = ?;", categorie)
	if err != nil {
		panic(err)
	}

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Likes, &post.Categorie); err != nil {
			panic(err)
		}
		post.UserName = GetName(post.UserID)

		posts = append(posts, post)
	}

	return posts
}

func GetUserId(Post_id int) int {
	post := GetPost(strconv.Itoa(Post_id))
	return post.UserID
}

// Fonction qui permet d'insérer un nouveau post avec l'id de l'utilisateur qui le poster,
// le titre du post, son commentaire ainsi que la catégprie de celui ci
func InsertPost(user_id int, title string, content string, categorie string) {
	db := connect_postsDB()

	query := `INSERT INTO posts (user_id, title,content,categorie) VALUES (?,?,?,?);`
	_, err := db.Exec(query, user_id, title, content, categorie)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion :", err)
		return
	}
	defer db.Close()
}

// Ajoute ou enlève 1 like sur un post
func NbLikes(id_post int, changement int) {
	db := connect_postsDB()
	defer db.Close()

	query := "UPDATE posts SET likes = likes + ? WHERE id = ?"
	_, err := db.Exec(query, changement, id_post)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du nombre de likes:", err)
	}

	db2 := connect_usersDB()
	defer db.Close()

	id := GetUserId(id_post)

	query2 := "UPDATE users SET nbLikes = nbLikes + ? WHERE id = ?"
	_, err2 := db2.Exec(query2, changement, id)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du nombre de likes:", err2)
	}
}

// Supprime un utilisateur en fonction de son id
func DeletePost(id string) {
	db := connect_postsDB()
	defer db.Close()

	query := `DELETE FROM posts WHERE id = ?;`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Erreur lors de la suppression :", err)
		return
	}
}
