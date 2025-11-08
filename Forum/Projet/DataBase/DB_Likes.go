package DataBase

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Structure qui reprend les caractéristiques de la Base de donnée
type Likes struct {
	ID      int
	Post_id int
	User_id int
	Status  int
}

// Fonction qui permet de retourner la connexion a la base de donnée posts
func connect_likesDB() *sql.DB {
	dbLikes, err := sql.Open("sqlite", "../DataBase/Stockage/Likes.db")
	if err != nil {
		panic(err)
	}

	// Vérifier la connexion
	if err := dbLikes.Ping(); err != nil {
		panic("Impossible de se connecter à la base de données dbLikes !")
	}

	return dbLikes
}

// Fonction qui créer la base de données posts si elle n'existe pas
// Si celle ci existe déjà alors elle reprend toutes les informations qui existes
func CreateTableLikes() {
	db := connect_likesDB()
	defer db.Close()

	// Structure de notre base de donnée
	query := `
	CREATE TABLE IF NOT EXISTS likes (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		status INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table 'likes' créée avec succès.")
}

// Fonction qui renvoie un tableau de tous les likes qui sont dans la base de donnée posts
func GetLikes() []Likes {
	db := connect_likesDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, post_id, user_id, status FROM likes;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Final []Likes
	for rows.Next() {
		var LikesFinal Likes

		if err := rows.Scan(&LikesFinal.ID, &LikesFinal.Post_id, &LikesFinal.User_id, &LikesFinal.Status); err != nil {
			panic(err)
		}

		Final = append(Final, LikesFinal)
	}
	return Final
}

func InsertLike(post_id int, user_id int) {
	db := connect_likesDB()
	defer db.Close()

	query := `INSERT INTO likes (post_id, user_id, status) VALUES (?,?,?);`
	_, err := db.Exec(query, post_id, user_id, 1)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion :", err)
		return
	}
	defer db.Close()
}

func InsertDislike(post_id int, user_id int) {
	db := connect_likesDB()
	defer db.Close()

	query := `INSERT INTO likes (post_id, user_id, status) VALUES (?,?,?);`
	_, err := db.Exec(query, post_id, user_id, -1)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion :", err)
		return
	}
	defer db.Close()
}

func DeleteLike(id string) {
	db := connect_likesDB()
	defer db.Close()

	query := `DELETE FROM likes WHERE id = ?;`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Erreur lors de la suppression :", err)
		return
	}
}

func DeleteLikes(id string) {
	db := connect_likesDB()
	defer db.Close()

	query := `DELETE FROM likes WHERE post_id = ?;`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Erreur lors de la suppression :", err)
		return
	}
}

func GetUserLikeStatus(postID int, userID int) int {
	db := connect_likesDB()
	defer db.Close()

	var status int
	err := db.QueryRow("SELECT status FROM likes WHERE post_id = ? AND user_id = ?", postID, userID).Scan(&status)
	if err != nil {
		return 0 // Aucun vote
	}
	return status
}
