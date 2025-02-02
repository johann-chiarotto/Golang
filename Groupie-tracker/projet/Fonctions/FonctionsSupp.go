package Fonctions

import (
	"fmt"
	"math/rand"
)

// Fonction utilisée seulement dans ce fichier
// Renvoie un booléen en fonction de si l'index passé en argument est dans le tableau passé en argument lui aussi
func InList(s int, tab []int) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] == s {
			return true
		}
	}
	return false
}

// Fonction qui renvoie un tableau avec des index aléatoires
// Les index vont de 0 à 4 et définissent l'odre des styles musicales proposés
func GenresAleatoire() []int {
	var tab []int

	for len(tab) < 5 {
		randomNumb := rand.Intn(5)
		if InList(randomNumb, tab) == false {
			tab = append(tab, randomNumb)
		}
	}
	return tab
}

// Fonction qui renvoie un style et un tableau comprennant des index de groupe de musiques
// Elle est appelée pour récupérer les informations de la fonction GenresAleatoire()
func Genres(index int) (string, []int) {
	//Liste des 5 genres disponibles
	TabStyles := []string{"Rock classique / Rock progressif", "Rap moderne / Hip-Hop contemporain ", "Electro-Pop ", "Pop traditionnelle / Pop mainstream", "Rock alternatif / Indie Rock "}
	//Initialisation et remplissage des tableaux en fonction des genres
	TabIndex := [][]int{
		{1, 3, 4, 9, 13, 14, 15, 16, 18, 19, 20, 49},
		{5, 24, 6, 7, 8, 25, 29, 30, 31, 43},
		{48, 52, 23},
		{11, 12, 47},
		{23, 37, 50},
	}

	var TabIndexFinal []int

	randomNumber := index

	//Logique de la fonction qui renvoie que 6 artistes si la liste est trop longue
	if len(TabIndex[randomNumber]) > 6 {
		for len(TabIndexFinal) < 6 {
			randomNumber2 := rand.Intn(len(TabIndex[randomNumber]))

			if InList(TabIndex[randomNumber][randomNumber2], TabIndexFinal) == false {
				TabIndexFinal = append(TabIndexFinal, TabIndex[randomNumber][randomNumber2])
			}
		}
		return TabStyles[randomNumber], TabIndexFinal
	}
	return TabStyles[randomNumber], TabIndex[randomNumber]
}

// Fonction qui renvoie aléatoirement un tableau d'index
func Recommandation() []int {
	var tab []int

	for len(tab) < 6 {
		random := rand.Intn(52)
		if InList(random+1, tab) == false {
			tab = append(tab, random+1)
		}
	}
	return tab
}

// Fonction qui renvoie le lien pour l'image en fonction de l'ID de la ville
// L'image est cherchée dans un fichier json créé pour cela
func GoodLink(id string) string {
	villes, err := RecupVilles()
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return "https://www.shutterstock.com/image-vector/404-error-icon-vector-symbol-260nw-1545236357.jpg"
	}

	for _, ville := range villes {
		if ville.ID == id {
			return ville.Image
		}
	}

	return "https://www.shutterstock.com/image-vector/404-error-icon-vector-symbol-260nw-1545236357.jpg"
}

// Fonction qui renvoie seulement le nom du pays sans la ville
func JustCity(s string) string {
	NewString := ""
	Passer := false

	for i := 0; i < len(s); i++ {
		if Passer {
			NewString += string(rune(s[i]))
		}
		if string(rune(s[i])) == "-" {
			Passer = true
			i++
		}
	}
	return NewString
}
