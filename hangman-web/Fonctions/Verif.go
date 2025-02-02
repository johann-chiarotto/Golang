package piscine

//Fonction qui fait le traitement de la proposition du joueur
//ici si il y a plus de 1 caractère
//s=proposition //i=instance  //mot=MotFinal  //try=Pour sauter le premier passage
func Mot(s string, i string, mot string) (string, int, int) {
	for k := 0; k < len(s); k++ {
		if s[k] < 64 || s[k] > 91 {
			return i, 0, 2
		}
	}

	mot = Upper(mot)
	i = Upper(i)
	if len(s) != len(mot) {
		return i, 2, 0
	} else if s == mot {
		return mot, 0, 0
	} else {
		return i, 2, 0
	}
}

//Fonction qui fait le traitement de la proposition du joueur
//ici si il y a 1 seul caractère
//s=proposition //i=instance  //mot=MotFinal  //LettresPasse=tableau de lettres déjà demmandé  //try=Pour sauter le premier passage
//renvoie : InstanceWordUp, NbViesUp, LetterAddList , num erreur
func Letter(s string, i string, mot string, LettresPasse []string) (string, int, string, int) {
	if s > string(rune(64)) && s < string(rune(91)) {
		if DejaDemmande(LettresPasse, s) {
			return i, 0, "", 1
		} else {
			if InTexte(mot, s) {
				MotACompleter := UpdateWord(i, mot, s)
				return MotACompleter, 0, s, 0
			} else {
				return i, 1, s, 0
			}
		}
	} else {
		return i, 0, "", 2
	}
}
