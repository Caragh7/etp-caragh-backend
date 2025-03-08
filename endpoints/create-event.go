package endpoints

import (
	"encoding/json"
	"net/http"
)

// Event represents an event with a title, description, venue, and date.
type Event struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Venue       string `json:"venue"`
	Date        string `json:"date"`
}

var eventsList []Event

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Event
	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	eventsList = append(eventsList, newEvent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(eventsList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
