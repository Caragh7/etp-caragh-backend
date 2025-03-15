package events

import (
	"encoding/json"
	"net/http"
)

// a single event
 type Event struct {
	Title string `json:"title"`
	Date string `json:"date"`
	VenueImageOne string `json:"venue_image_one"`
	Description string `json:"description"`
	Tickets int `json:"tickets"`
	Genre string `json:"genre"`
	TicketPrice int `json:"ticketPrice"`
  }
  
  
  // the initial state
  var events = []Event{
	  {
		Title: "Emergency Intercom Live Show",
		Date: "2025-06-15",
		VenueImageOne: "/etp-caragh/images/3arena_1.jpg",
		Description: "A comedy podcast hosted by Enya and Drew",
		Tickets: 700,
		Genre: "Comedy",
		TicketPrice: 30.0,
	  },
	  {
		Title: "Metallica",
		Date: "2025-07-20",
		VenueImageOne: "/etp-caragh/images/croke-park-1.jpg",
		Description: "A once in a lifetime experience world at Croke Park",
		Tickets: 80000,
		Genre: "Metal",
		TicketPrice: 140.0,
	  },
	  {
		Title: "Taylor Swift: Eras Tour",
		Date: "2025-08-10",
		VenueImageOne: "/etp-caragh/images/o2_1.jpg",
		Description: "Experience Taylor Swift performing all her eras live!",
		Tickets: 75000,
		Genre: "Pop",
		TicketPrice: 200.0,
	  },
	  {
		Title: "The Weeknd: After Hours Tour",
		Date: "2025-09-05",
		VenueImageOne: "/etp-caragh/images/thomond_2.jpeg",
		Description:
		  "A cinematic experience with The Weeknd and his biggest hits",
		Tickets: 60000,
		Genre: "Hip-Hop",
		TicketPrice: 90.0,
	  },
	  {
		Title: "Yung Lean World Tour",
		Date: "2025-10-12",
		VenueImageOne: "/etp-caragh/images/slane_2.jpg",
		Description:
		  "Join Yung Lean for a once in a lifetime experience in Slane Castle",
		Tickets: 50000,
		Genre: "Rap",
		TicketPrice: 120.0,
	  },
	  {
		Title: "Billie Eilish: Happier Than Ever Tour",
		Date: "2025-11-20",
		VenueImageOne: "/etp-caragh/images/croke-park-2.jpg",
		Description: "A unique and intimate experience with Billie Eilish",
		Tickets: 30000,
		Genre: "Pop",
		TicketPrice: 175.0,
	  },
	  {
		Title: "Noel Miller: Company Lot",
		Date: "2025-12-01",
		VenueImageOne: "/etp-caragh/images/ul_2.jpg",
		Description:
		  "Join Noel Miller in his live comedy show for a belly of laughs",
		Tickets: 700,
		Genre: "Comedy",
		TicketPrice: 50.0,
	  },
  };

func GetEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}