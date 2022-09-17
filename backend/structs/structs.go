package structs

// LiftEntry Standard struct that we'll use for storing lift data
type LiftEntry struct {
	event      string
	date       string
	gender     string
	name       string
	bodyweight float32
	sn1        float32
	sn2        float32
	sn3        float32
	cj1        float32
	cj2        float32
	cj3        float32
	bestSn     float32
	bestCJ     float32
	total      float32
	sinclair   float32
	federation string
}

// APIPayload Standard struct that we'll use for storing lift data
type APIPayload struct {
	Event      string  `json:"event"`
	Date       string  `json:"date"`
	Gender     string  `json:"gender"`
	Name       string  `json:"name"`
	Bodyweight float32 `json:"bodyweight"`
	Sn1        float32 `json:"snatch_1"`
	Sn2        float32 `json:"snatch_2"`
	Sn3        float32 `json:"snatch_3"`
	CJ1        float32 `json:"cj_1"`
	CJ2        float32 `json:"cj_2"`
	CJ3        float32 `json:"cj_3"`
	BestSn     float32 `json:"best_sn"`
	BestCJ     float32 `json:"best_cj"`
	Total      float32 `json:"total"`
	Sinclair   float32 `json:"sinclair"`
	Federation string  `json:"federation"`
}
