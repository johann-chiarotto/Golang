package main

import (
	"fmt"
	hgc "hangmanweb/Fonctions"
	"html/template"
	"net/http"
)

// varaibles d'initialisation utilent pour la partie
var Status bool
var FinalWord string
var InstanceWord string
var NbVies int
var LettresPasse []string
var Message string
var Replay bool // Nouvelle variable pour gérer le redémarrage

// Structure utile pour l'envoie des informations a la page HTML
type Bloc struct {
	FinalWord     string
	InstanceWord  string
	NbVies        int
	MessageSupp   string
	Lettres       string
	Hangman       string
	Back          string
	Border        string
	Border1       string
	Input         string
	Replay        bool // Ajouter cette donnée au modèle
	MessageButton string
}

// fonction main
func main() {
	initializeGame() // Initialisation de la première partie

	// Configure le système de fichiers pour les fichiers statiques
	fs := http.FileServer(http.Dir("./CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))

	image := http.FileServer(http.Dir("../Images"))
	http.Handle("/Images/", http.StripPrefix("/Images/", image))

	// Une page d'accueil pour vérifier que le serveur fonctionne
	http.HandleFunc("/", HomePage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}

// Réinitialise les variables pour une nouvelle partie si le joueur veux rejouer
func initializeGame() {
	Status = true
	FinalWord = hgc.Random()
	FinalWord = hgc.Upper(FinalWord)
	InstanceWord = hgc.LetterReveal(FinalWord)
	NbVies = 10
	LettresPasse = []string{}
	Message = ""
	Replay = false
}

// Fonction qui s'occupe des GET et des POST
// Cette fonction fait le lien entre les requêtes entre le front end et le back end
func HomePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		HandleGet(w, r)

	case "POST":
		HandlePost(w, r)
	}
}

// Partie qui s'occupe du traitement front end vers baack end
// Les datas sont envoyés grâce a la stucture que l'on a créer
func HandleGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))

	var data Bloc
	//Si le joueur a encore des vies et si il n'a pas encore trouvé le mot
	if NbVies > 0 && Status {
		data = Bloc{
			FinalWord:     FinalWord,
			InstanceWord:  InstanceWord,
			NbVies:        NbVies,
			MessageSupp:   Message,
			Lettres:       hgc.AffichageLettresPasse(LettresPasse),
			Hangman:       hgc.Affichage(10 - NbVies),
			Back:          hgc.Back(10 - NbVies),
			Border:        hgc.Border(10 - NbVies),
			Border1:       hgc.Border1(10 - NbVies),
			Input:         hgc.Input(10 - NbVies),
			MessageButton: "Submit",
		}
		//Si le joueur a encore des vies mais a trouvé le mot
	} else if NbVies > 0 && Status == false {
		Replay = true
		data = Bloc{
			FinalWord:     "",
			InstanceWord:  InstanceWord,
			NbVies:        NbVies,
			MessageSupp:   "Congrats! YOU WON!!!",
			Lettres:       hgc.AffichageLettresPasse(LettresPasse),
			Hangman:       hgc.Affichage(10 - NbVies),
			Back:          hgc.Back(10 - NbVies),
			Border:        hgc.Border(10 - NbVies),
			Border1:       hgc.Border1(10 - NbVies),
			Input:         hgc.Input(10 - NbVies),
			MessageButton: "Next",
		}
		//Si le joueur a perdu et qu'il n'a fonc plus de vie et il n'a pas trouver le mot
	} else {
		Replay = true // Permet d'afficher une option pour rejouer
		data = Bloc{
			FinalWord:    FinalWord,
			InstanceWord: FinalWord,
			NbVies:       0,
			MessageSupp:  "You LOST... The word was : ",
			Lettres:      hgc.AffichageLettresPasse(LettresPasse),
			Hangman:      hgc.Affichage(10),
			Back:         hgc.Back(10),
			Border:       hgc.Border(10),
			Border1:      hgc.Border1(10),
			Input:        hgc.Input(10),
			Replay:       Replay,
		}
	}

	//Envoie les données sur la page web
	tmpl.Execute(w, data)
}

// Partie qui s'occupe du traitement back end vers front end
// Les datas sont envoyés grâce a la stucture que l'on a créer
func HandlePost(w http.ResponseWriter, r *http.Request) {
	Message = ""
	// Traite les données envoyées via le formulaire
	if err := r.ParseForm(); err != nil {
		fmt.Printf("Erreur ParseForm() : %v\n", err)
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	// Si le joueur souhaite rejouer
	if Replay {
		initializeGame()
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// récupère la valeur et la met en majuscule
	variable := hgc.Upper(r.Form.Get("input"))

	// Logique principale
	if Status {
		//variables utilent seulement dans la manche et réinitialisé a chaque début de manches
		var InstanceWordUp string
		var NbViesUp int
		var LetterAddList string
		var erreur int

		if NbVies > 0 {
			//Si la proposition a plus de 1 caractère
			if len(variable) > 1 {
				InstanceWordUp, NbViesUp, erreur = hgc.Mot(variable, InstanceWord, FinalWord)
			} else { //Si la proposition a 1 seul caractère
				//récupère la nouvelle version du mot, le nombre de vies que le joueur perd, le lettre a ajouter a la liste et un numero suivant si il y a une erreur ou non
				InstanceWordUp, NbViesUp, LetterAddList, erreur = hgc.Letter(variable, InstanceWord, FinalWord, LettresPasse)
			}

			//Update des variables avec les nouvelles versions
			InstanceWord = InstanceWordUp
			NbVies -= NbViesUp
			LettresPasse = append(LettresPasse, LetterAddList)

			//Regarde si le mot est complètement dévoiler ou non
			if hgc.IsFinish(InstanceWord) {
				Message = "Congrats! YOU WON!!! Play again?"
				Status = false
			}

			//Gestion des erreurs
			if erreur == 1 {
				Message = "You've already typed this letter"
			} else if erreur == 2 {
				Message = "You must only type letters"
			}
		}
	}

	//Renvoie les informations vers la page HMLT
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
