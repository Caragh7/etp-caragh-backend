package venues 

import (
	"encoding/json"
	"net/http"
)

func CreateVenue(w http.ResponseWriter, r *http.Request) {
	var venue Venue
	_ = json.NewDecoder(r.Body).Decode(&venue)

	venues = append(venues, venue)
	json.NewEncoder(w).Encode(venue)
}