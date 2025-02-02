package Fonctions

import (
	"html/template"
	"net/http"
	"strconv"
)

// Structure de base que l'on va envoyer au front-end
type Data struct {
	Artistes        []Artiste
	ArtisteSolo     Artiste
	ConcertsAffiche []Concerts
	Recommandation  []Artiste
	Genres          []TypeGenre
}

// //////
// PAGE//
// //////
// Fonction qui gère le lien back-end/front-end de la HomePage
func HomePage(w http.ResponseWriter, r *http.Request) {
	artistes, err := RecupArtistes()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./homepage.html"))

	//Recommandations
	var TabRecommmandations []Artiste
	TabRecommandationsID := Recommandation()

	for i := 0; i < 6; i++ {
		var ArtisteInstance Artiste

		for _, artist := range artistes { //Enregistre toutes les données de l'artiste au moment donner
			if strconv.Itoa(artist.ID) == strconv.Itoa(TabRecommandationsID[i]) {
				ArtisteInstance = artist
			}
		}
		TabRecommmandations = append(TabRecommmandations, ArtisteInstance)
	}

	//Genres
	var GenresFinal []TypeGenre

	TabGenresAleatoire := GenresAleatoire()

	//Pour les 5 genres
	for j := 0; j < len(TabGenresAleatoire); j++ {
		var Instance TypeGenre
		var Style string
		var TabIndex []int

		//Recupère le style et le tableau des artistes à afficher grâce à leurs ID
		Style, TabIndex = Genres(TabGenresAleatoire[j])

		Instance.Name = Style

		//Boucle sur la longueur de la liste des artistes à afficher
		var TableauArtistesGenre []Artiste
		for i := 0; i < len(TabIndex); i++ {
			var ArtisteInstance Artiste

			for _, artist := range artistes { //Enregistre toutes les données de l'artiste à l'instant T
				if strconv.Itoa(artist.ID) == strconv.Itoa(TabIndex[i]) {
					ArtisteInstance = artist
				}
			}
			//Rentre l'artiste dans le tableau avec son index
			TableauArtistesGenre = append(TableauArtistesGenre, ArtisteInstance)
		}
		Instance.ListeArtistes = TableauArtistesGenre

		GenresFinal = append(GenresFinal, Instance)
	}

	data := Data{
		Artistes:       artistes,
		Genres:         GenresFinal,
		Recommandation: TabRecommmandations,
	}

	tmpl.Execute(w, data)
}

// Fonction qui gère le lien back-end/front-end de la page Artistes
// Elle recupère tous les artistes de l'API et envoie des datas
func ArtistesPage(w http.ResponseWriter, r *http.Request) {
	artistes, err := RecupArtistes()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./artistes.html"))

	data := Data{
		Artistes: artistes,
	}
	tmpl.Execute(w, data)
}

// Fonction qui gère le lien back-end/front-end de la page Artiste
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Récupère le paramètre `id` depuis l'URL

	artistes, err := RecupArtistes()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("./artiste.html"))

	var Artiste Artiste

	//Recupère le bon artiste
	for _, artist := range artistes {
		if strconv.Itoa(artist.ID) == id {
			Artiste = artist
		}
	}

	// Créer le bon affichage du string des membres exemple : nom1, nom2, nom3
	for i := 0; i < len(Artiste.Members)-1; i++ {
		Artiste.MembersS += Artiste.Members[i] + `, `
	}
	Artiste.MembersS += Artiste.Members[len(Artiste.Members)-1]

	data := Data{
		ArtisteSolo: Artiste,
	}
	tmpl.Execute(w, data)
}

// Fonction qui gère le lien back-end/front-end de la page Concert
func ConcertPage(w http.ResponseWriter, r *http.Request) {
	artistes, err := RecupArtistes()
	relations, err := RecupRelations()

	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	id := r.URL.Query().Get("id") // Récupère le paramètre `id` depuis l'URL
	tmpl := template.Must(template.ParseFiles("./concert.html"))

	var Artiste Artiste
	// Récupère les données des artistes
	for _, artist := range artistes {
		if strconv.Itoa(artist.ID) == id {
			Artiste = artist
		}
	}

	var Final []Concerts

	//Rcupère toutes les informations sur le bon artiste
	for _, relation := range relations {
		if strconv.Itoa(relation.ID) == id {
			locations := relation.DatesLocations

			for location, dates := range locations {
				var ConcertsAffiche Concerts

				ConcertsAffiche.ID = relation.ID

				locationUpdate := ""
				for i := 0; i < len(location); i++ {
					if string(rune(location[i])) == "_" {
						locationUpdate += " "
					} else if string(rune(location[i])) == "-" {
						locationUpdate += " - "
					} else {
						locationUpdate += string(rune(location[i]))
					}
				}

				ConcertsAffiche.Lieu = locationUpdate
				ConcertsAffiche.Link = GoodLink(JustCity(ConcertsAffiche.Lieu))
				ConcertsAffiche.Date = dates

				Final = append(Final, ConcertsAffiche)
			}
		}
	}

	data := Data{
		ArtisteSolo:     Artiste,
		ConcertsAffiche: Final,
	}
	tmpl.Execute(w, data)
}
