package events

import (
	"encoding/json"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	_ = json.NewDecoder(r.Body).Decode(&event)
	events = append(events, event)
	json.NewEncoder(w).Encode(event)
}
