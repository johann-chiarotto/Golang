package piscine

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// je pense que le probleme vien du readfile donc faut essayer avec bufio mais jsp comment faire donc faut regarder pour faire ca
//apres tous le reste doit être adaptable sans trop de changements

func PrintWordLigne(word string) {
	//récupère toutes les données du fichier texte
	LettersListeBase, err := ioutil.ReadFile("Letters.txt")
	// si le contenu du fichier est nul (vide) on renvoi une erreur
	if err != nil {
		log.Fatal(err)
	}
	//sépare lignes par lignes
	LettersListe := strings.Split(string(LettersListeBase), "\n")

	//enlève la retour a la ligne a la fin de chaque lignes
	for a := 0; a < len(LettersListe); a++ {
		mot := LettersListe[a]
		var motRune []byte
		for i := 0; i < len(mot)-1; i++ {
			motRune = append(motRune, mot[i])
		}
		StringF := ""
		for j := 0; j < len(motRune); j++ {
			StringF += string(rune(motRune[j]))
		}
		LettersListe[a] = StringF
	}

	//transforme le mot en tableau de runes
	var WordInRune []rune
	for i := 0; i < len(word); i++ {
		WordInRune = append(WordInRune, rune(word[i]))
	}
	var tab [95][9]string
	for i := 0; i < 95; i++ {
		for j := 0; j < 9; j++ {
			tab[i][j] = LettersListe[i*9+(j)]
		}
	}

	//affiche la mot en ascii art
	for a := 0; a < 9; a++ {
		for b := 0; b < len(WordInRune); b++ {
			fmt.Print(tab[WordInRune[b]-32][a])
		}
		fmt.Println()
	}
}
