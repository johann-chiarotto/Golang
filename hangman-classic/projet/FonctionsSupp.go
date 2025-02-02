package piscine

import (
	"fmt"
	"strings"
)

// fonction qui prend en paramètre un texte et le met en majuscule
func Upper(lettre string) string {
	if len(lettre) == 1 && lettre > string(rune(96)) && lettre < string(rune(123)) {
		lettre2 := strings.ToUpper(lettre)
		lettre = lettre2
	}
	return lettre
}

// fonction qui regarde grâce a une liste et une lettre donner en paramètre si celle ci à déja était demmander auparavant ou non
func DejaDemmander(LettresPasser []string, lettre string) bool {
	for i := 0; i < len(LettresPasser); i++ {
		if lettre == LettresPasser[i] {
			return true
		}
	}
	return false
}

// fonction qui sert a regarder si une lettre est présente dans le texte ou non
// similaire a celle d'avant mais regarde dans notre texte car propbleme si plusieurs fois la meme lettre mais qu'une seule révélée au début
func InTexte(motRune []rune, lettre string) bool {
	for i := 0; i < len(motRune); i++ {
		if lettre == string(motRune[i]) {
			return true
		}
	}
	return false
}

// fonction qui transforme au(x) bon(s) endroit(s) le caractère "_" en la bonne lettre et renvoie une chaine de caractère à jour
func UpdateWord(MotACompleter string, motRune []rune, lettre string) string {
	//trouver tous les endroits de la où ya les lettres
	//remplace dans le mot les bonnes lettres (anciennement "_")
	var MotAcompleterRune []rune
	for i := 0; i < len(MotACompleter); i++ {
		MotAcompleterRune = append(MotAcompleterRune, rune(MotACompleter[i]))
	}

	MotACompleterInstance := ""
	for i := 0; i < len(motRune); i++ {
		if lettre == string(motRune[i]) {
			MotACompleterInstance += string(motRune[i])
		} else if MotAcompleterRune[i] != '_' {
			MotACompleterInstance += string(MotAcompleterRune[i])
		} else {
			MotACompleterInstance += "_"
		}
	}
	return MotACompleterInstance
}

// focntion qui regarde si le joueur à trouver toutes les lettres ou non
// si oui la fonction renvoie un booleen true, sinon elle renvoie un booleen false
func IsFinish(MotACompleter string) bool {
	//reagrde si le mot est trouver en entier ou non
	//transforme la chaine de caractère en tableau de rune
	var MotAcompleterRune []rune
	for i := 0; i < len(MotACompleter); i++ {
		MotAcompleterRune = append(MotAcompleterRune, rune(MotACompleter[i]))
	}
	//regarde si il reste des '_', synonyme de lettre non trouver
	for i := 0; i < len(MotACompleter); i++ {
		if MotAcompleterRune[i] == '_' {
			return false
		}
	}
	return true
}

// Affiche un message d'erreur pour prevenir qu'il nous reste tans de vies et affiche l'état du pendu
func MessageError1(message1 string, nb_vie int, message2 string) {
	fmt.Println()
	fmt.Printf(message1)
	fmt.Printf(string(rune(nb_vie + 48)))
	fmt.Printf(message2)
	fmt.Println()
	Affichage(10 - nb_vie)
	fmt.Println()
}

// Fait l'affichage de la fin du jeu
func MessageFin(motRune []rune, motInstance string, lettre string) bool {
	mot2 := ""
	for i := 0; i < len(motRune); i++ {
		mot2 += string(motRune[i])
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Congrats !")
	return false
}
