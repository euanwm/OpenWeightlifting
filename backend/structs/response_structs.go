package structs

type EventResponse struct {
	Event Event  `json:"event"`
	Lifts []Lift `json:"lifts"`
}
