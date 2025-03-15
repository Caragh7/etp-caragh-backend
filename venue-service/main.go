package main

import (
	venues "etp-caragh-backend/venue-service/routes"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/venues", venues.GetVenues)
	http.HandleFunc("/venues/create", venues.CreateVenue)

	log.Println("Venue Service started on port 8082")
	http.ListenAndServe(":8082", nil)
}
