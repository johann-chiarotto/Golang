package DataBase

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// Structure qui reprend les caractéristiques de la Base de donnée
type Comment struct {
	ID       int
	UserID   int
	PostID   int
	Content  string
	UserName string
	UserPDP  string
}

// Fonction qui permet de retourner la connexion a la base de donnée comments
func connect_commentsDB() *sql.DB {
	dbComments, err := sql.Open("sqlite", "../DataBase/Stockage/Comments.db")
	if err != nil {
		panic(err)
	}

	// Vérifier la connexion
	if err := dbComments.Ping(); err != nil {
		panic("Impossible de se connecter à la base de données dbComments !")
	}
	return dbComments
}

// Fonction qui créer la base de données comments si elle n'existe pas
// Si celle ci existe déjà alors elle reprend toutes les informations qui existes
func CreateTableComments() {
	db := connect_commentsDB()
	defer db.Close()

	// Structure de notre base de donnée
	query := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
	);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Table 'comments' créée avec succès.")
}

// Fonction qui renvoie un tableau de tous les comments qui sont dans la base de donnée posts
func GetComments() []Comment {
	var Final []Comment

	db := connect_commentsDB()

	rows, err := db.Query("SELECT id, user_id, post_id, content FROM posts;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var CommentFinal Comment

		var id int
		var user_id int
		var post_id int
		var content string
		if err := rows.Scan(&id, &user_id, &post_id, &content); err != nil {
			panic(err)
		}

		CommentFinal.ID = id
		CommentFinal.UserID = user_id
		CommentFinal.PostID = post_id
		CommentFinal.Content = content

		Final = append(Final, CommentFinal)

	}
	return Final
}

// Fonction qui renvoie un tableau de tous les commentaires qui on pour caractéristique le meme id du post
// Alors cette fonction renvoie un tableau de tous les commenataires d'un post passer en paramètres grâce a son id
func GetCommentsPost(id string) []Comment {
	db := connect_commentsDB()

	var comments []Comment

	rows, err := db.Query("SELECT id, user_id, post_id, content FROM comments WHERE post_id = ?", id)
	if err != nil {
		log.Println("Erreur lors de la récupération des commentaires:", err)
		return comments
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content); err != nil {
			log.Println("Erreur lors du scan des commentaires:", err)
			continue
		}
		comment.UserName = GetName(comment.UserID)
		comment.UserPDP = GetPdp(comment.UserID)
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		log.Println("Erreur après l'itération des résultats:", err)
	}

	return comments
}

// Fonction qui permet d'insérer un nouveau comment avec l'id de l'utilisateur qui le poster,
// l'id du post ou le commentaire à était publier et le contenu de ce commentaire
func InsertComment(user_id int, post_id string, content string) {
	db := connect_commentsDB()

	query := `INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?);`
	_, err := db.Exec(query, user_id, post_id, content)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion :", err)
		return
	}
	defer db.Close()
}

func DeleteComment(commentId string) {
	db := connect_commentsDB()

	_, err := db.Exec("DELETE FROM comments WHERE id = ?", commentId)
	if err != nil {
		log.Println(err)
	}
}
