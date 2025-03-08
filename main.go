package main

import (
	"etp-caragh/endpoints"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/venues", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			endpoints.GetVenues(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			endpoints.CreateEvent(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
