package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.ArtistsHandler)
	http.HandleFunc("/info/", handlers.InfoHandler)
	http.HandleFunc("/locations/", handlers.LocationsHandler)
	http.HandleFunc("/dates/", handlers.DatesHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))

	port := ":1945"
	fmt.Printf("http://localhost%s\n", port)

	http.ListenAndServe(port, nil)
}
