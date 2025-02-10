package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter is missing", http.StatusBadRequest)
		return
	}

	artists, err := GetArtists()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching artists: %v", err), http.StatusInternalServerError)
		return
	}

	var results []map[string]string
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, map[string]string{
				"label": artist.Name,
				"type":  "artist/band",
				"url":   fmt.Sprintf("/info/%d", artist.ID),
			})
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				results = append(results, map[string]string{
					"label": member,
					"type":  "member",
					"url":   fmt.Sprintf("/info/%d", artist.ID),
				})
			}
		}
		if strings.Contains(strings.ToLower(artist.Locations), strings.ToLower(query)) {
			results = append(results, map[string]string{
				"label": artist.Locations,
				"type":  "location",
				"url":   fmt.Sprintf("/info/%d", artist.ID),
			})
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) {
			results = append(results, map[string]string{
				"label": artist.FirstAlbum,
				"type":  "first album",
				"url":   fmt.Sprintf("/info/%d", artist.ID),
			})
		}
		if strings.Contains(fmt.Sprintf("%d", artist.CreationDate), query) {
			results = append(results, map[string]string{
				"label": fmt.Sprintf("Année de création: %d", artist.CreationDate),
				"type":  "creation date",
				"url":   fmt.Sprintf("/info/%d", artist.ID),
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
