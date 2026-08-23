package structs

type WeightClass struct {
	Gender   string
	Upper    WeightKg
	Lower    WeightKg
	DateFrom string
	DateTo   string
}

type ContainerTime struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
	Sec  int `json:"sec"`
}

type NameSearchResults struct {
	// todo: refactor this so we don't have to worry about case sensitivity on the items within the slice
	Names []NameSearch `json:"names"`
	Total int          `json:"total"`
}

type NameSimilarity struct {
	NameStr    string  `json:"name"`
	Federation string  `json:"federation"`
	Score      float32 `json:"score"`
}

type NameSimilarityResults struct {
	Names []NameSimilarity `json:"names"`
	Total int              `json:"total"`
}

type RivalsResult struct {
	Rivals []Rival `json:"rivals"`
	Total  int     `json:"total"`
}

type Rival struct {
	Position   int      `json:"position"`
	Total      WeightKg `json:"total"`
	Lifter     string   `json:"lifter"`
	Federation string   `json:"federation"`
}

type RivalsCombined struct {
	FederationRivals RivalsResult `json:"federationrivals"`
	CombinedRivals   RivalsResult `json:"combinedrivals"`
}

type NameSearch struct {
	NameStr    string `json:"name"`
	Gender     string `json:"gender"`
	Federation string `json:"federation"`
}

type LifterHistory struct {
	NameStr string  `json:"name"`
	Lifts   []*Lift `json:"lifts"`
}

type LifterStats struct {
	BestSnatch       WeightKg `json:"best_snatch"`
	BestCJ           WeightKg `json:"best_cj"`
	BestTotal        WeightKg `json:"best_total"`
	MakeRateSnatches []int    `json:"make_rate_snatches"`
	MakeRateCJ       []int    `json:"make_rate_cj"`
}

type LeaderboardData struct {
	AllTotals    []*Lift
	AllSinclairs []*Lift
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

type SearchLeaderboardRequest struct {
	ActiveQuery LeaderboardPayload `json:"active_query"`
	LifterData  NameSearch         `json:"lifter_data"`
}

type SearchLeaderboardResult struct {
	LifterData NameSearch         `json:"lifter_data"`
	Position   int                `json:"position"`
	Query      LeaderboardPayload `json:"query"`
}

type EventsData struct {
	Events []*Event `json:"events"`
}

type Event struct {
	Name       string  `json:"name"`
	Date       string  `json:"date"`
	Federation string  `json:"federation"`
	CSVID      string  `json:"id"`
	Results    []*Lift `json:"-"`
}

type LifterRoster struct {
	Lifters []*Lifter `json:"lifters"`
	index   map[string]*Lifter
}

type Lifter struct {
	Gender            string  `json:"gender"`
	Name              string  `json:"name"`
	currentAge        uint8   // todo: implement age calculation & linking
	PrimaryFederation string  `json:"federation"`
	Lifts             []*Lift `json:"-"` // back-reference; would cycle through Lift.Lifter
}

type AllLifts struct {
	Lifts []*Lift `json:"lifts"`
}

type Lift struct {
	Event      *Event   `json:"event"` // parent; already the context when nested under Event.Results/EventResponse
	Lifter     *Lifter  `json:"lifter"`
	ageOnDay   uint8    // todo: implement age calculation & linking
	Category   string   `json:"category"`
	Bodyweight WeightKg `json:"bodyweight"`
	Sn1        WeightKg `json:"snatch_1"`
	Sn2        WeightKg `json:"snatch_2"`
	Sn3        WeightKg `json:"snatch_3"`
	CJ1        WeightKg `json:"cj_1"`
	CJ2        WeightKg `json:"cj_2"`
	CJ3        WeightKg `json:"cj_3"`
	BestSn     WeightKg `json:"best_snatch"`
	BestCJ     WeightKg `json:"best_cj"`
	Total      WeightKg `json:"total"`
	Sinclair   float64  `json:"sinclair"` // todo: change this to a key:value so we can differentiate between qpoints, sinclair etc.
}

type LeaderboardResponse struct {
	Size int     `json:"size"`
	Data []*Lift `json:"data"`
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
	Events []Event `json:"events"`
}
