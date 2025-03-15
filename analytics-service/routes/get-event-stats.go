package analytics

import (
	"encoding/json"
	"net/http"
)

type AnalyticsData struct {
	Metric string `json:"metric"`
	Value  int    `json:"value"`
}

func GetEventAnalytics(w http.ResponseWriter, r *http.Request) {
	data := []AnalyticsData{
		{"Total Events", 25},
		{"Tickets Sold", 300},
	}
	json.NewEncoder(w).Encode(data)
}