package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/backend/structures"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func GetArtistByID(id int) (structures.Artist, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)

	response, err := http.Get(url)
	if err != nil {
		return structures.Artist{}, err
	}
	defer response.Body.Close()

	var artist structures.Artist
	if err := json.NewDecoder(response.Body).Decode(&artist); err != nil {
		return structures.Artist{}, err
	}

	return artist, nil
}

func GetTopTracks(artistName string) ([]structures.Track, error) {
	url := fmt.Sprintf("https://api.deezer.com/search?q=%s&limit=3", strings.ReplaceAll(artistName, " ", "%20"))

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var deezerResponse structures.DeezerResponse
	if err := json.NewDecoder(response.Body).Decode(&deezerResponse); err != nil {
		return nil, err
	}

	return deezerResponse.Data, nil
}

func GetRelations(artistID int) (structures.Relation, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", artistID)

	response, err := http.Get(url)
	if err != nil {
		return structures.Relation{}, err
	}
	defer response.Body.Close()

	var relation structures.Relation
	if err := json.NewDecoder(response.Body).Decode(&relation); err != nil {
		return structures.Relation{}, err
	}

	return relation, nil
}

func capitalizeWords(s string) string {
	return strings.Title(s)
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, err := GetArtistByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching artist info: %v", err), http.StatusInternalServerError)
		return
	}

	artist.FirstAlbum = strings.ReplaceAll(artist.FirstAlbum, "-", ".")

	topTracks, err := GetTopTracks(artist.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching top tracks: %v", err), http.StatusInternalServerError)
		return
	}

	relation, err := GetRelations(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching relations: %v", err), http.StatusInternalServerError)
		return
	}

	updatedDatesLocations := make(map[string][]string)
	for location, dates := range relation.DatesLocations {
		for i, date := range dates {
			dates[i] = strings.ReplaceAll(date, "-", ".")
		}
		updatedDatesLocations[location] = dates
	}

	formattedDatesLocations := make(map[string][]string)
	for location, dates := range updatedDatesLocations {
		formattedLocation := strings.ReplaceAll(location, "_", " ")
		formattedLocation = strings.ReplaceAll(formattedLocation, "-", ", ")
		formattedLocation = capitalizeWords(formattedLocation)
		formattedDatesLocations[formattedLocation] = dates
	}

	relation.DatesLocations = formattedDatesLocations

	data := struct {
		Artist    structures.Artist
		TopTracks []structures.Track
		Relation  structures.Relation
	}{
		Artist:    artist,
		TopTracks: topTracks,
		Relation:  relation,
	}

	tmpl, err := template.ParseFiles("../frontend/templates/info.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
