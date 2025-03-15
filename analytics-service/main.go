package main

import (
	analytics "etp-caragh-backend/analytics-service/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/analytics/events", analytics.GetEventAnalytics)
	http.HandleFunc("/analytics/venues", analytics.GetVenueAnalytics)

	log.Println("Analytics Service started on port 8083")
	http.ListenAndServe(":8083", nil)
}