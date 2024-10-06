package structs

type WeightClass struct {
	Gender string
	Upper  float32
	Lower  float32
}

type ContainerTime struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
	Sec  int `json:"sec"`
}

type AllData struct {
	Lifts []Entry
}

type NameSearchResults struct {
	// todo: refactor this so we don't have to worry about case sensitivity on the items within the slice
	Names []struct {
		Name       string
		Federation string
	} `json:"names"`
	Total int `json:"total"`
}

type NameSearch struct {
	NameStr    string `json:"name"`
	Federation string `json:"federation"`
}

type ChartData struct {
	Dates   []string       `json:"labels"`
	SubData []ChartSubData `json:"datasets"`
}

type ChartSubData struct {
	Title     string    `json:"label"`
	DataSlice []float32 `json:"data"`
}

type LifterHistory struct {
	NameStr string      `json:"name"`
	Lifts   []Entry     `json:"lifts"`
	Graph   ChartData   `json:"graph"`
	Stats   LifterStats `json:"stats"`
}

type LifterStats struct {
	BestSnatch       float32 `json:"best_snatch"`
	BestCJ           float32 `json:"best_cj"`
	BestTotal        float32 `json:"best_total"`
	MakeRateSnatches []int   `json:"make_rate_snatches"`
	MakeRateCJ       []int   `json:"make_rate_cj"`
}

type LeaderboardData struct {
	AllTotals    []Entry
	AllSinclairs []Entry
}

// LeaderboardPayload Incoming request payload
type LeaderboardPayload struct {
	Start       int    `form:"start"`
	Stop        int    `form:"stop"`
	SortBy      string `form:"sortby"`
	Federation  string `form:"federation"`
	WeightClass string `form:"weightclass"`
	Year        string `form:"year"`
	StartDate   string `form:"startdate"`
	EndDate     string `form:"enddate"`
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

type LeaderboardResponse struct {
	Size int     `json:"size"`
	Data []Entry `json:"data"`
}

// EventsMetaData Internal struct for storing event metadata
type EventsMetaData struct {
	Name       []string
	Federation []string
	Date       []string
	ID         []string
}

type SingleEventMetaData struct {
	Name       string `json:"name"`
	Federation string `json:"federation"`
	Date       string `json:"date"`
	ID         string `json:"id"`
}

type EventSearch struct {
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type SingleEvent struct {
	Federation string `json:"federation"`
	ID         string `json:"id"`
}

type EventsList struct {
	Events []SingleEventMetaData `json:"events"`
}

type LiftReport struct {
	ReportedLift Entry  `json:"lift"`
	Comments     string `json:"comments"`
}

type BeanCounter struct {
	Bytes uint64
}
