package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/backend/structures"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func GetLocationByID(id int) (structures.Location, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)

	response, err := http.Get(url)
	if err != nil {
		return structures.Location{}, err
	}
	defer response.Body.Close()

	var location structures.Location
	if err := json.NewDecoder(response.Body).Decode(&location); err != nil {
		return structures.Location{}, err
	}

	artist, err := GetArtistByID(id)
	if err != nil {
		return structures.Location{}, err
	}

	location.Image = artist.Image
	location.Name = artist.Name

	return location, nil
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	intID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %v", err), http.StatusBadRequest)
		return
	}

	location, err := GetLocationByID(intID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching location: %v", err), http.StatusInternalServerError)
		return
	}

	if location.ID == 0 {
		http.Error(w, "No location found for this ID.", http.StatusNotFound)
		return
	}

	for i, loc := range location.Locations {
		loc = strings.ReplaceAll(strings.ReplaceAll(loc, "-", ", "), "_", " ")
		location.Locations[i] = strings.Title(loc)
	}

	tmpl, err := template.ParseFiles("../frontend/templates/locations.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, location)
}
