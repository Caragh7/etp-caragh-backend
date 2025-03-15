package venues

import (
	"encoding/json"
	"net/http"
)

type Venue struct {
	Name string `json:"name"`
	Location string `json:"location"`
	Capacity int `json:"capacity"`
	Image_one string `json:"image_one"`
	Image_two string `json:"image_two"`
	Description string `json:"description"`
}

var venues = []Venue{
	{
		Name: "Croke Park",
		Location: "Dublin",
		Capacity: 80000,
		Image_one: "/etp-caragh/images/croke-park-1.jpg",
		Image_two: "/etp-caragh/images/croke-park-2.jpg",
		Description:
		  "One of Ireland’s largest stadiums, located in Dublin, famous for hosting Gaelic games, concerts, and major sporting events",
	  },
	  {
		Name: "Thomond Park",
		Location: "Limerick",
		Capacity: 26000,
		Image_one: "/etp-caragh/images/thomond_1.jpg",
		Image_two: "/etp-caragh/images/thomond_2.jpeg",
		Description:
		  "A renowned rugby stadium in Limerick, home to Munster Rugby, known for its electric atmosphere and passionate supporters",
	  },
	  {
		Name: "3 Arena",
		Location: "Dublin",
		Capacity: 13500,
		Image_one: "/etp-caragh/images/3arena_1.jpg",
		Image_two: "/etp-caragh/images/3arena_2.jpg",
		Description:
		  "A modern indoor venue in Dublin, hosting international concerts, comedy shows, and live performances with state-of-the-art acoustics",
	  },
	  {
		Name: "O2 Arena",
		Location: "London",
		Capacity: 20000,
		Image_one: "/etp-caragh/images/o2_2.jpg",
		Image_two: "/etp-caragh/images/o2_1.jpg",
		Description:
		  " A world-class entertainment venue in London, featuring concerts, sporting events, and cultural performances in a massive dome-shaped arena",
	  },
	  {
		Name: "Slane Castle",
		Location: "Meath",
		Capacity: 8000,
		Image_one: "/etp-caragh/images/slane_1.jpg",
		Image_two: "/etp-caragh/images/slane_2.jpg",
		Description:
		  "A historic venue in Meath, famous for its stunning setting and legendary outdoor concerts featuring global music icons",
	  },
	  {
		Name: "University Concert Hall",
		Location: "Limerick",
		Capacity: 1000,
		Image_one: "/etp-caragh/images/ul_1.jpg",
		Image_two: "/etp-caragh/images/ul_2.jpg",
		Description:
		  "A cultural landmark in Limerick, offering an intimate setting for classical music, theater, and live performances",
	  },
	}



func GetVenues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(venues); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
