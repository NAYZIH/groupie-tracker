package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))

	http.HandleFunc("/", handlers.ArtistsHandler)
	http.HandleFunc("/info/", handlers.InfoHandler)
	http.HandleFunc("/locations/", handlers.LocationsHandler)
	http.HandleFunc("/dates/", handlers.DatesHandler)

	port := ":1945"
	fmt.Printf("http://localhost%s\n", port)

	http.ListenAndServe(port, nil)
}
