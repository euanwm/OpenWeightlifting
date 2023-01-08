package enum

import "backend/structs"

const (
	Male     string = "male"
	Female   string = "female"
	Unknown  string = "unknown"
	Total    string = "total"
	Sinclair string = "sinclair"
	//BWL - British Weightlifting
	BWL string = "UK"
	//USAW - United States of America Weightlifting
	USAW string = "US"
	//IWF - Internation Weightlifting Federation
	IWF string = "IWF"
	//ALLFEDS - Pretty self-explanatory
	ALLFEDS string = "allfeds"
	//ALLWEIGHTS - Yes
	ALLCATS string = "allcats"
	//AWF - Australian Weightlifting Federation
	AWF               string  = "AUS"
	MaxTotal          float32 = 510
	MinimumBodyweight float32 = 20
	MaximumBodyweight float32 = 300
)

var DefaultPayload = structs.LeaderboardPayload{
	Start:       0,
	Stop:        500,
	Gender:      "male",
	SortBy:      "total",
	Federation:  "allfeds",
	WeightClass: "allcats",
}

var WeightClassList = map[string]structs.WeightClass{
	"M55":   {Upper: 55.00, Lower: 0},
	"M61":   {Upper: 61.00, Lower: 55.01},
	"M67":   {Upper: 67.00, Lower: 61.01},
	"M73":   {Upper: 73.00, Lower: 67.01},
	"M81":   {Upper: 81.00, Lower: 73.01},
	"M89":   {Upper: 89.00, Lower: 81.01},
	"M96":   {Upper: 96.00, Lower: 89.01},
	"M102":  {Upper: 102.00, Lower: 96.01},
	"M109":  {Upper: 109.00, Lower: 102.01},
	"M109+": {Upper: MaximumBodyweight, Lower: 109.01},
	"F45":   {Upper: 45.00, Lower: 0},
	"F49":   {Upper: 49.00, Lower: 45.01},
	"F55":   {Upper: 55.00, Lower: 49.01},
	"F59":   {Upper: 59.00, Lower: 55.01},
	"F64":   {Upper: 64.00, Lower: 59.01},
	"F71":   {Upper: 71.00, Lower: 64.01},
	"F76":   {Upper: 76.00, Lower: 71.01},
	"F81":   {Upper: 81.00, Lower: 76.01},
	"F87":   {Upper: 87.00, Lower: 81.01},
	"F87+":  {Upper: MaximumBodyweight, Lower: 87.01},
}
