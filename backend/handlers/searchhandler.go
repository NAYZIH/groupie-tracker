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
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
