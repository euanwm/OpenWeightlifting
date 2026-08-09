package structs

type FullEvent struct {
	Event Event  `json:"event"`
	Lifts []Lift `json:"lifts"`
}

type LiftResponse struct {
}
