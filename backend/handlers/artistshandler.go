package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/backend/structures"
	"html/template"
	"net/http"
)

func GetArtists() ([]structures.Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []structures.Artist
	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := GetArtists()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching artists: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("../frontend/templates/artists.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)
}
