package piscine

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// nous permet de choisir quel hangman afficher suivant le nombre d'essai (a partir d'un fichier texte où sont stockés tout nos hangman)
func Affichage(etape int) {

	//affiche l'état du pendu en fonction du nombre d'erreurs que le joueur a déja effectuer
	var tab []string
	// va mettre dans les variables Man et err le contenu du fichier texte AffichageHangman
	Man, err := ioutil.ReadFile("AffichageHangman.txt")
	// si le contenu du fichier est nul (vide) on renvoi une erreur
	if err != nil {
		log.Fatal(err)
	}
	// sépare chaque ligne dans le fichier texte
	LignesMan := strings.Split(string(Man), "\n")

	// définit le bloque que l'on va garder (7 en hauteur) suivant notre valeur d'essai
	for i := 7*etape - 7 + 1*etape - 2; i < 7*etape+1*etape-2; i++ {
		tab = append(tab, LignesMan[i+1])
	}

	// là où on stocke notre hangman à afficher dans une seule string
	caracteres := ""
	for j := 0; j < 7; j++ {
		caracteres += tab[j]
		caracteres += string('\n')
	}
	fmt.Println(caracteres)
}
