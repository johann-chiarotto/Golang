package piscine

import (
	"strings"
)

// fonction qui prend en paramètre un texte et le met en majuscule
func Upper(lettre string) string {
	lettreUp := ""
	for i := 0; i < len(lettre); i++ {
		lettreUp += strings.ToUpper(string(rune(lettre[i])))
	}
	return lettreUp
}

// fonction qui regarde grâce a une liste et une lettre donner en paramètre si celle ci à déja était demmander auparavant ou non
func DejaDemmande(LettresPasse []string, lettre string) bool {
	for i := 0; i < len(LettresPasse); i++ {
		if lettre == LettresPasse[i] {
			return true
		}
	}
	return false
}

// fonction qui sert a regarder si une lettre est présente dans le texte ou non
// similaire a celle d'avant mais regarde dans notre texte car propbleme si plusieurs fois la meme lettre mais qu'une seule révélée au début
func InTexte(mot string, lettre string) bool {
	var motRune []rune

	for i := 0; i < len(mot); i++ {
		motRune = append(motRune, rune(mot[i]))
	}

	for i := 0; i < len(motRune); i++ {
		if lettre == string(motRune[i]) {
			return true
		}
	}
	return false
}

// fonction qui transforme au(x) bon(s) endroit(s) le caractère "_" en la bonne lettre et renvoie une chaine de caractère à jour
func UpdateWord(MotACompleter string, mot string, lettre string) string {
	var motRune []rune

	for i := 0; i < len(mot); i++ {
		motRune = append(motRune, rune(mot[i]))
	}

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

// Fonction qui fait un string des lettres passées pour l'affichage web
func AffichageLettresPasse(tab []string) string {
	s := ""

	for i := 0; i < len(tab); i++ {
		if tab[i] != "" {
			s += tab[i] + " "
		}
	}
	return s
}
