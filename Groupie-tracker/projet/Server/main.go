package main

import (
	"fmt"
	"net/http"
	fc "tracker/Fonctions"
)

func main() {

	fs := http.FileServer(http.Dir("./CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	http.HandleFunc("/", fc.HomePage)
	http.HandleFunc("/Artistes", fc.ArtistesPage)
	http.HandleFunc("/Artiste", fc.ArtistPage)
	http.HandleFunc("/Concert", fc.ConcertPage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
