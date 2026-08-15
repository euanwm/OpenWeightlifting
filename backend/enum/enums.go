package enum

import "strings"

// ClassifyGender reduces a raw CSV gender/category string (which may carry
// weight-class/age info too, e.g. "Women Juniors 59kg") down to a stable
// Male/Female/Unknown value suitable for identifying a lifter across their career.
func ClassifyGender(raw string) string {
	switch {
	case raw == Male:
		return Male
	case raw == Female:
		return Female
	case strings.Contains(raw, "Women") || strings.Contains(raw, "female"):
		return Female
	case strings.Contains(raw, "Men") || strings.Contains(raw, "male"): // todo: this is temporary probably
		return Male
	default:
		return Unknown
	}
}

const (
	Male     string = "male"
	Female   string = "female"
	Unknown  string = "unknown"
	Total    string = "total"
	Sinclair string = "sinclair"
	// ALLFEDS - Pretty self-explanatory
	ALLFEDS string = "allfeds"
	// ALLCATS ALLWEIGHTS - Yes
	ALLCATS     string = "allcats"
	AllYearsStr string = "69"
	AllYears    int    = 69
	ZeroDate    string = "0001-01-01"
	MaxDate     string = "2100-00-00"
	// lift rules
	MaxSnatch         float64 = 240
	MaxCleanAndJerk   float64 = 280
	MaxTotal          float64 = 510
	MinimumBodyweight float64 = 20
	MaximumBodyweight float64 = 300
	BestSnatch        string  = "BestSn"
	BestCJ            string  = "BestCJ"
	Bodyweight        string  = "Bodyweight"
	Snatch            string  = "snatch"
	CleanAndJerk      string  = "clean_and_jerk"
)
