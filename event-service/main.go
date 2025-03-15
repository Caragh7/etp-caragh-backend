package main

import (
	events "etp-caragh-backend/event-service/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/events", events.GetEvents)
	http.HandleFunc("/events/create", events.CreateEvent)

	log.Println("Event Service started on port 8081")
	http.ListenAndServe(":8081", nil)
}