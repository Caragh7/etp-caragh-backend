package analytics

import (
	"encoding/json"
	"net/http"
)


func GetVenueAnalytics(w http.ResponseWriter, r *http.Request) {
	data := []AnalyticsData{
		{"Total Venues", 5},
		{"Most Popular Venue", 1},
	}
	json.NewEncoder(w).Encode(data)
}