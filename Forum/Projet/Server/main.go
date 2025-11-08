package main

import (
	fc "forum/Fonctions"
	"log"
	"net/http"
)

func main() {
	fc.SetDB()
	fc.SetHttp()

	//inst.InsertAdmin("Admin", "forum1234")

	log.Println("\nServeur démarré sur https://localhost:8080\n")
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur HTTPS : ", err)
	}
}
