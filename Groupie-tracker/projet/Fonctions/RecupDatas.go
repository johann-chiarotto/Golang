package Fonctions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Déclaration de toutes les structures utilisées dans le projet
type Artiste struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	MembersS     string
}

type Concerts struct {
	ID   int
	Lieu string
	Date []string
	Link string
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type LocationIndex struct {
	Index []Location `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DateIndex struct {
	Index []Date `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

type Ville struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}

type TypeGenre struct {
	ListeArtistes []Artiste
	Name          string
}

// /////
// API// Récupère toutes les données des API et du fichier json des villes
// /////
func RecupArtistes() ([]Artiste, error) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"

	// Requête HTTP GET
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Décoder la réponse JSON
	var artistes []Artiste
	if err := json.NewDecoder(resp.Body).Decode(&artistes); err != nil {
		return nil, err
	}

	return artistes, nil
}

func RecupLocations() ([]Location, error) {
	apiURL := "https://groupietrackers.herokuapp.com/api/locations"

	// Requête HTTP GET avec timeout
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête GET : %w", err)
	}
	defer resp.Body.Close()

	// Vérifier le statut HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statut HTTP inattendu : %s", resp.Status)
	}

	// Décoder la réponse JSON
	var locationIndex LocationIndex
	if err := json.NewDecoder(resp.Body).Decode(&locationIndex); err != nil {
		return nil, fmt.Errorf("erreur lors de la désérialisation JSON : %w", err)
	}

	// Retourner les données
	return locationIndex.Index, nil
}

func RecupDates() ([]Date, error) {
	apiURL := "https://groupietrackers.herokuapp.com/api/dates"

	// Requête HTTP GET avec timeout
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête GET : %w", err)
	}
	defer resp.Body.Close()

	// Vérifier le statut HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statut HTTP inattendu : %s", resp.Status)
	}

	// Décoder la réponse JSON
	var dateIndex DateIndex
	if err := json.NewDecoder(resp.Body).Decode(&dateIndex); err != nil {
		return nil, fmt.Errorf("erreur lors de la désérialisation JSON : %w", err)
	}

	// Retourner les données
	return dateIndex.Index, nil
}

func RecupRelations() ([]Relation, error) {
	apiURL := "https://groupietrackers.herokuapp.com/api/relation"

	// Requête HTTP GET avec timeout
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête GET : %w", err)
	}
	defer resp.Body.Close()

	// Vérifier le statut HTTP
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statut HTTP inattendu : %s", resp.Status)
	}

	// Décoder la réponse JSON
	var relationIndex RelationIndex
	if err := json.NewDecoder(resp.Body).Decode(&relationIndex); err != nil {
		return nil, fmt.Errorf("erreur lors de la désérialisation JSON : %w", err)
	}

	// Retourner les données
	return relationIndex.Index, nil
}

func RecupVilles() (map[string]Ville, error) {
	fichier := "../Fonctions/locations_with_images.json"

	// Ouvrir le fichier
	file, err := os.Open(fichier)
	if err != nil {
		return nil, fmt.Errorf("échec de l'ouverture du fichier: %w", err)
	}
	defer file.Close()

	// Décoder le contenu JSON
	var villes map[string]Ville
	if err := json.NewDecoder(file).Decode(&villes); err != nil {
		return nil, fmt.Errorf("échec du décodage JSON: %w", err)
	}

	return villes, nil
}
