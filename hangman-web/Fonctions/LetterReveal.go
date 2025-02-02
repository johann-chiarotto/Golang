package piscine

import (
	"math/rand"
	"sort"
	"strings"
)

func LetterReveal(word string) string {
	//cherche les indices de premières lettres qui vont apparaîtrent dès le début de la partie

	//met en majuscule le mot passer en paramètre
	wordSec := strings.ToUpper(word)
	word = wordSec
	//créer un tableau qui contiendras les indices des lettres à faire apparaître au début de la partie
	var Index []int
	fin := (len(word))/2 - 1
	running := true
	sauter := 0
	//boucle qui continue de tourner tans que l'on a pas l'indices ou les indices différents de nos lettres à faire apparaître
	for running {
		//si l'on a pas encore l'ensemble de nos indices l'algorithme continue de générer un indice au hasard et s'il n'est pas encore utiliser on le garde
		if fin > 0 {
			IndexSupp := rand.Intn(len(word) - 1)
			for j := 0; j < len(Index); j++ {
				if IndexSupp == Index[j] {
					sauter++
				}
			}
			//si l'index n'était pas utiliser alors on l'ajoute au tableau d'indices
			if sauter == 0 {
				Index = append(Index, IndexSupp)
				fin--
			}
			//si l'on a déjà l'ensemble de nos indices, on arrete la boucle qui génère nos indices au hasard
		} else {
			running = false
		}
		sauter = 0
	}
	//créer une tableau de runes avec les premières lettres visibles
	var NewWord []rune
	//transoforme notre mot en un tableau de runes
	WordInRune := []rune(word)
	//ajoute a notre tableau "NewWord", pour le moment vide, le même nombre de "_" que le nombre de lettres de notre mot
	for i := 0; i < len(word); i++ {
		NewWord = append(NewWord, '_')
	}
	//tri notre tableau d'indices dans l'odre croissant
	sort.Ints(Index)
	//remplace les _ par les bonnes letres à l'endroit où les lettres doivent êtres montrer
	avancement := 0
	for i := 0; i < len(word); i++ {
		if avancement < len(Index) && i == Index[avancement] {
			NewWord[i] = WordInRune[i]
			avancement++
		}
	}
	//change le mot qui est sous la forme d'un tableau de rune vers une string
	StringWord := ""

	for i := 0; i < len(WordInRune); i++ {
		StringWord += string(NewWord[i])
	}
	//renvoie la string du mot non compléter mais avec les premières lettres affichées
	return StringWord
}
