package main

import (
	"fmt"
	"piscine"
	"strings"
)

// fonction principale
func main() {
	runningP := true
	nb_reussites := 0
	for runningP {
		//mot sous la forme entier en string sans \n
		mot := piscine.Random()
		mot2 := ""
		for i := 0; i < len(mot)-1; i++ {
			mot2 += string(rune(mot[i]))
		}
		mot3 := strings.ToUpper(mot2)
		mot = mot3

		//mot entier sous la forme tableau de runes
		var motRune []rune
		for i := 0; i < len(mot); i++ {
			motRune = append(motRune, rune(mot[i]))
		}

		//mot qui change a chaque fois sous la forme d'une string
		MotACompleter := piscine.LetterReveal(mot)

		running := true
		nb_vie := 10
		var LettresPasser []string

		fmt.Println()
		fmt.Println("Good Luck, you have 10 attempts.")
		fmt.Println()

		for running {
			motInstance := ""
			for i := 0; i < len(MotACompleter); i++ {
				motInstance += string(rune(MotACompleter[i])) + " "
			}
			//affiche le mot a compléter
			piscine.PrintWordLigne(MotACompleter)
			fmt.Println()
			//demmande une lettre
			var lettre string
			fmt.Printf("Letters already typed : ")
			if len(LettresPasser) > 0 {
				fmt.Printf(LettresPasser[0])
			}
			for i := 1; i < len(LettresPasser); i++ {
				fmt.Printf(", ")
				fmt.Printf(LettresPasser[i])
			}

			fmt.Println()
			fmt.Printf("CHOOSE : ")
			for lettre == "" {
				fmt.Scanf("%s", &lettre)
			}
			lettreInstance := strings.ToUpper(lettre)
			lettre = lettreInstance
			//met en majuscule le mot ou la lettre que le joueur demmande
			lettre = piscine.Upper(lettre)
			//soit la longueur est supérieur à un seul caractère
			if len(lettre) > 1 {
				// 4 EasterEggs
				if lettre == "QUIT" {
					running = false
					runningP = false
					fmt.Println()
					fmt.Println("You just quitted the game !")
				} else if lettre == "BLEH!" {
					piscine.ReadFileAndDisplay("blehcatascii.txt")
				} else if lettre == "BAPTISTE" {
					piscine.ReadFileAndDisplay("EasterEggBaptiste.txt")
				} else if lettre == "ADRIEN" {
					piscine.ReadFileAndDisplay("EasterEggAdrien.txt")
				} else if lettre == "JOHANN" {
					piscine.ReadFileAndDisplay("EasterEggJohann.txt")
				} else if lettre == "VICENZZO" {
					piscine.ReadFileAndDisplay("EasterEggVicenzzo.txt")
				} else if len(lettre) != len(mot) { //regarde si la longueur du mot donner par le joueur est égal a la longeur du mot a deviner
					if nb_vie < 2 {
						nb_vie -= 1
					} else {
						nb_vie -= 2
					}
					piscine.MessageError1("Wrong ! This is not the right word, you have ", nb_vie, " attempts left")
					//regarde si le mot donner par le joueur est le même que celui a deviner
				} else if lettre == mot {
					running = piscine.MessageFin(motRune, motInstance, lettre)
					nb_reussites++
					//sinon c'est que le joueur n'a pas donner le bon mot et alors il perd 2 vies
				} else {
					if nb_vie < 2 {
						nb_vie -= 1
					} else {
						nb_vie -= 2
					}
					piscine.MessageError1("Wrong ! This is not the right word, you have ", nb_vie, " attempts left")
				}
			} else {
				//regarde si le lettre est la bonne quand la longueur du texte du joueur est égale a 1
				//si le caractère est une lettre (c'est forcement une lettre majuscule car changé avant)
				if lettre > string(rune(64)) && lettre < string(rune(91)) {
					//si la lettre a déja était demander auparavent
					if piscine.DejaDemmander(LettresPasser, lettre) {
						fmt.Println("You've already typed this letter.")
					} else {
						//si la lettre a déja était dévoiler par le début du jeu
						if piscine.InTexte(motRune, lettre) {
							fmt.Println()
							//remet a jour le mot a deviner avec la nouvelle lettre apparente
							MotACompleter = piscine.UpdateWord(MotACompleter, motRune, lettre)
							//ajoute dans une liste toutes les lettres que le joueur a demmander depuis le début
							LettresPasser = append(LettresPasser, lettre)

						} else {
							//si la lettre n'est pas présente dans le mot
							//renvoie d'un message et suppression d'une vie au joueur
							nb_vie--
							LettresPasser = append(LettresPasser, lettre)
							piscine.MessageError1("Wrong ! This letter is not in the word, you have ", nb_vie, " attempts left")
						}
					}
				} else {
					//prévient que le joueuer peux seulement donner des lettres pour deviner le mot
					fmt.Println("You must only type letters.")
				}

			}
			//si le programe trouve que le mot a était trouvé, il lance l'affichage de la fin et arrête le programme
			if piscine.IsFinish(MotACompleter) {
				running = piscine.MessageFin(motRune, motInstance, lettre)
				nb_reussites++
			}
			//regarde si il reste assez de vies au joueur pour continuer a jouer encore au moins une fois,
			//si il n'a pas assez de vie il lui affiche comme quoi il a perdu
			if nb_vie < 1 {
				running = false
				fmt.Println("You LOSE...")
				runningP = false
			}
		}
		fmt.Println()
		fmt.Println()
		fmt.Println("The word was : ")
		fmt.Println()
		piscine.PrintWordLigne(mot)

		if runningP {
			var reponse string
			fmt.Println()
			fmt.Printf("WOULD YOU LIKE TO CONTINUE ? ( Y / N ) : ")
			for reponse == "" {
				fmt.Scanf("%s", &reponse)
			}
			reponseInstance := strings.ToUpper(reponse)
			reponse = reponseInstance

			if reponse == "N" || reponse == "NO" {
				runningP = false
			}

		} else if runningP == false {
			if nb_reussites > 1 {
				fmt.Println()
				fmt.Println("You won", nb_reussites, "times in a row.")
			} else {
				fmt.Println()
				fmt.Println("You won", nb_reussites, "time in a row.")
			}
		}
	}
}
