package piscine

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

// Fonction Random qui retourne un mot aléatoire à partir d'un fichier texte
func Random() string {
	//ouvre le fichier texte "word.txt" où se situt tous les mots
	content, err := ioutil.ReadFile("../Fonctions/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	//modifie le tableau pour faire une liste de tous les mots séparer par un retour à la ligne
	ListeWords := strings.Split(string(content), "\n")

	//choisi un mot au hasard dans la liste de mot grâce a la fonction random
	n := rand.Intn(len(ListeWords) - 1)

	//renvoie le mot choisi au hasard sous la forme d'un string
	word := ListeWords[n]
	wordF := ""
	for i := 0; i < len(word)-1; i++ {
		wordF += string(rune(word[i]))
	}
	return wordF
}
