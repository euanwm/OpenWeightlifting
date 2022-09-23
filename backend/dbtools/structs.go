package dbtools

type LeaderboardData struct {
	MaleTotals   []Entry
	FemaleTotals []Entry
}

// LeaderboardPayload Incoming request payload
type LeaderboardPayload struct {
	Start  int    `json:"start"`
	Stop   int    `json:"stop"`
	Gender string `json:"gender"`
}

// Entry Standard struct that we'll use for storing raw lift data
type Entry struct {
	Event      string  `json:"event"`
	Date       string  `json:"date"`
	Gender     string  `json:"gender"`
	Name       string  `json:"lifter_name"`
	Bodyweight string  `json:"bodyweight"`
	Sn1        string  `json:"snatch_1"`
	Sn2        string  `json:"snatch_2"`
	Sn3        string  `json:"snatch_3"`
	CJ1        string  `json:"cj_1"`
	CJ2        string  `json:"cj_2"`
	CJ3        string  `json:"cj_3"`
	BestSn     string  `json:"best_snatch"`
	BestCJ     string  `json:"best_cj"`
	Total      float32 `json:"total"`
	Federation string  `json:"country"`
}
