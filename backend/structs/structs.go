package structs

type TestPayload struct {
	Hour int
	Min  int
	Sec  int
}

type AllData struct {
	Lifts []Entry
}

type NameSearchResults struct {
	Names []string `json:"names"`
}

type NameSearch struct {
	NameStr string
}

type LifterHistory struct {
	NameStr string  `json:"name"`
	Lifts   []Entry `json:"lifts"`
}

type LeaderboardData struct {
	AllNames        []string
	AllData         []Entry
	MaleTotals      []Entry
	FemaleTotals    []Entry
	MaleSinclairs   []Entry
	FemaleSinclairs []Entry
}

// LeaderboardPayload Incoming request payload
type LeaderboardPayload struct {
	Start      int    `json:"start"`
	Stop       int    `json:"stop"`
	Gender     string `json:"gender"`
	SortBy     string `json:"sortby"`
	Federation string `json:"federation"`
}

// Entry Standard structs that we'll use for storing raw lift data
type Entry struct {
	Event      string  `json:"event"`
	Date       string  `json:"date"`
	Gender     string  `json:"gender"`
	Name       string  `json:"lifter_name"`
	Bodyweight float32 `json:"bodyweight"`
	Sn1        float32 `json:"snatch_1"`
	Sn2        float32 `json:"snatch_2"`
	Sn3        float32 `json:"snatch_3"`
	CJ1        float32 `json:"cj_1"`
	CJ2        float32 `json:"cj_2"`
	CJ3        float32 `json:"cj_3"`
	BestSn     float32 `json:"best_snatch"`
	BestCJ     float32 `json:"best_cj"`
	Total      float32 `json:"total"`
	Sinclair   float32 `json:"sinclair"`
	Federation string  `json:"country"`
}
