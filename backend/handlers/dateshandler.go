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

func GetDatesByID(id int) (structures.Dates, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)

	response, err := http.Get(url)
	if err != nil {
		return structures.Dates{}, err
	}
	defer response.Body.Close()

	var dates structures.Dates
	if err := json.NewDecoder(response.Body).Decode(&dates); err != nil {
		return structures.Dates{}, err
	}

	artist, err := GetArtistByID(id)
	if err != nil {
		return structures.Dates{}, err
	}

	dates.Image = artist.Image
	dates.Name = artist.Name

	for i, date := range dates.Dates {
		dates.Dates[i] = strings.ReplaceAll(strings.TrimPrefix(date, "*"), "-", ".")
	}

	return dates, nil
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
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

	dates, err := GetDatesByID(intID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching dates: %v", err), http.StatusInternalServerError)
		return
	}

	if dates.ID == 0 {
		http.Error(w, "No dates found for this ID.", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("../frontend/templates/dates.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, dates)
}
