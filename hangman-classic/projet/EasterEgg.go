package piscine

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadFileByLine lit un fichier texte ligne par ligne et les affiche
func ReadFileAndDisplay(filename string) {
	// Ouvre le fichier
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err) // Affiche une erreur et arrête le programme si le fichier n'est pas trouvé
	}
	defer file.Close()

	// Créer un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	// Lire chaque ligne et l'afficher
	for scanner.Scan() {
		line := scanner.Text() // Met chaque ligne dans la variable `line`
		fmt.Println(line)      // Affiche la ligne
	}

	// Vérifie s'il y a des erreurs lors de la lecture du fichier
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
