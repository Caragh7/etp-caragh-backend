package endpoints

import (
	"encoding/json"
	"net/http"
)

type Venue struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Seats    int    `json:"size"`
}

var venues = []Venue{
	{Name: "Croke Park", Location: "Dublin City", Seats: 1000},
	{Name: "Slane Castle", Location: "Slane", Seats: 1500},
}

func GetVenues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(venues); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
