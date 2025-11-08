package DataBase

import (
	"database/sql"
	"fmt"
	ch "forum/Chiffrement"
	"log"
	"strconv"

	_ "modernc.org/sqlite"
)

// Structure qui reprend les caractéristiques de la Base de donnée
type User struct {
	Id         int
	Username   string
	Password   string
	Photo      string
	NbPosts    int
	NbLikes    int
	NbComments int
	Created_at string
	Admin      int
}

// Fonction qui permet de retourner la connexion a la base de donnée users
func connect_usersDB() *sql.DB {
	dbUsers, err := sql.Open("sqlite", "../DataBase/Stockage/Users.db")
	if err != nil {
		panic(err)
	}

	// Vérifier la connexion
	if err := dbUsers.Ping(); err != nil {
		panic("Impossible de se connecter à la base de données dbUsers !")
	}
	return dbUsers
}

// Fonction qui créer la base de données users si elle n'existe pas
// Si celle ci existe déjà alors elle reprend toutes les informations qui existes
func CreateTableUsers() {
	db := connect_usersDB()
	defer db.Close()

	// Structure de notre base de donnée
	query := ` 
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		photo TEXT DEFAULT "https://www.shutterstock.com/image-vector/default-avatar-profile-icon-social-600nw-1906669723.jpg",
		nbPosts INTEGER DEFAULT 0,
		nbLikes INTEGER DEFAULT 0,
		nbComments INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		admin INTEGER DEFAULT 0
	);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table 'users' créée avec succès.")
}

// Récupère l'ensemble de la base de donnée users sous la forme d'un tableau de User
func GetUsers() []User {
	db := connect_usersDB()
	defer db.Close()

	var users []User
	rows, err := db.Query("SELECT id, username, password, photo, nbPosts, nbLikes, nbComments, created_at, admin FROM users;")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Photo, &user.NbPosts, &user.NbLikes, &user.NbComments, &user.Created_at, &user.Admin); err != nil {
			panic(err)
		}

		users = append(users, user)
	}
	return users
}

// Fonction qui renvoie un User en fonction d'un ID passer en paramètre
func GetUser(id string) User {
	db := connect_usersDB()
	defer db.Close()

	var user User
	query := `SELECT id, username, password, photo, nbPosts, nbLikes, nbComments, created_at, admin FROM users WHERE id = ?;`

	err := db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Password, &user.Photo, &user.NbPosts, &user.NbLikes, &user.NbComments, &user.Created_at, &user.Admin)
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'utilisateur :", err)
		return User{}
	}

	return user
}

// Fonction qui renvoie le lien de la photo de profil d'un utilisateur en fonction se son ID prit en paramètres
func GetPdp(id int) string {
	db := connect_usersDB()
	var S string

	query := "SELECT photo FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&S)
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'utilisateur :", err)
		return "https://www.shutterstock.com/image-vector/default-avatar-profile-icon-social-600nw-1906669723.jpg"
	}
	return S
}

// Fonction qui renvoie le Username correpondant a l'id donner en paramètre
func GetName(UserID int) string {
	User := GetUser(strconv.Itoa(UserID))
	if User.Username == "" {
		return "Iconnu"
	}
	return User.Username
}

// Fonction qui permet d'insérer un nouveau user avec son pseudo, son mdp et un indicateur d'admin
// Le mdp doit arriver en clair ( car chiffrer dans la fonction )
func InsertUser(username, password string, admin int) {
	db := connect_usersDB()

	mdp, err1 := ch.HashPassword(password) //Chiffre le mdp
	if err1 != nil {
		log.Fatal(err1)
	}

	query := `INSERT INTO users (username, password, admin) VALUES (?, ?, ?);`
	_, err := db.Exec(query, username, mdp, admin)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion :", err)
		return
	}
}

// Fonction accessible seulement dans le back pour rajouter un user admin sans que ce soit possible depuis l'interface utilisateur ( sécurité )
func InsertAdmin(username, password string) {
	InsertUser(username, password, 1)
}

// Supprime un utilisateur en fonction de son id
func DeleteUser(id string) {
	db := connect_usersDB()
	defer db.Close()

	query := `DELETE FROM users WHERE id = ?;`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Erreur lors de la suppression :", err)
		return
	}
}

// Fonction qui change le mdp avec en paramètres le nouveau mdp et d'id de l'utilisateur
// La vérification des accords de changement de mdp doivent-êtres effectuer en dehors de cette fonction
func ChangePassword(newPassword string, userID int) {
	db := connect_usersDB()
	defer db.Close()

	query := "UPDATE users SET password = ? WHERE id = ?"
	_, err := db.Exec(query, newPassword, userID)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du mot de passe :", err)
	}
}

// Fonction qui change la photo de profil avec en paramètres la nouvelle pdp et l'id de l'utilisateur
// La vérification des accords de changement de pdp doivent-êtres effectuer en dehors de cette fonction
func ChangePdp(pdp string, userID int) {
	db := connect_usersDB()
	defer db.Close()

	query := "UPDATE users SET photo = ? WHERE id = ?"
	_, err := db.Exec(query, pdp, userID)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour de votre photo de profil :", err)
	}
}

// Ajoute 1 au nombre de posts que le base de donnée a l'user qui a créer un Post
func NbPostUp(id int) {
	db := connect_usersDB()
	defer db.Close()

	user := GetUser(strconv.Itoa(id))
	resultat := user.NbPosts
	resultat++

	query := "UPDATE users SET nbPosts = ? WHERE id = ?"
	_, err := db.Exec(query, resultat, id)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du nombre de posts:", err)
	}
}

// Ajoute 1 au nombre de comments que le base de donnée a l'user qui a créer un Comment
func NbCommentsUp(id int) {
	db := connect_usersDB()
	defer db.Close()

	user := GetUser(strconv.Itoa(id))
	resultat := user.NbComments
	resultat++

	query := "UPDATE users SET nbComments = ? WHERE id = ?"
	_, err := db.Exec(query, resultat, id)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du nombre de comments:", err)
	}
}
