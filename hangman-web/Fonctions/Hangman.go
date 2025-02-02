package piscine

//Fonction qui donne le chemin vers la bonne image suivant l'état du pendu
func Affichage(etape int) string {
	tab := []string{"../Images/hangman0.png", "../Images/hangman1.png", "../Images/hangman2.png", "../Images/hangman3.png", "../Images/hangman4.png", "../Images/hangman5.png", "../Images/hangman6.png", "../Images/hangman7.png", "../Images/hangman8.png", "../Images/hangman9.png", "../Images/hangman10.png"}
	return tab[etape]
}
