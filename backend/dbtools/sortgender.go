package dbtools

import (
	"backend/enum"
	"backend/sinclair"
	"backend/structs"
	"backend/utilities"
	"strings"
)

// SortGender Splits results into 3 categories, male, female, and unknown.
func SortGender(bigData [][]string) (male structs.AllData, female structs.AllData, unknown structs.AllData) {
	for _, contents := range bigData {
		dataStruct := assignStruct(contents)
		gender := setGender(&dataStruct)
		switch gender {
		case enum.Male:
			if dataStruct.Total > 0 && dataStruct.Total < enum.MaxTotal && dataStruct.Bodyweight > enum.MinimumBodyweight {
				sinclair.CalcSinclair(&dataStruct, true)
			}
			male.Lifts = append(male.Lifts, dataStruct)
		case enum.Female:
			if dataStruct.Total > 0 && dataStruct.Total < enum.MaxTotal && dataStruct.Bodyweight > enum.MinimumBodyweight {
				sinclair.CalcSinclair(&dataStruct, false)
			}
			female.Lifts = append(female.Lifts, dataStruct)
		case enum.Unknown:
			unknown.Lifts = append(unknown.Lifts, dataStruct)
		}
	}
	return
}

func setGender(entry *structs.Entry) (gender string) {
	if entry.Gender == enum.Male {
		return enum.Male
	} else if entry.Gender == enum.Female {
		return enum.Female
	} else if strings.Contains(entry.Gender, "Men") {
		return enum.Male
	} else if strings.Contains(entry.Gender, "Women") {
		return enum.Female
	} else {
		return enum.Unknown
	}
}

func assignStruct(line []string) (lineStruct structs.Entry) {
	lineStruct = structs.Entry{
		Event:      line[0],
		Date:       line[1],
		Gender:     line[2],
		Name:       line[3],
		Bodyweight: utilities.Float(line[4]),
		Sn1:        utilities.Float(line[5]),
		Sn2:        utilities.Float(line[6]),
		Sn3:        utilities.Float(line[7]),
		CJ1:        utilities.Float(line[8]),
		CJ2:        utilities.Float(line[9]),
		CJ3:        utilities.Float(line[10]),
		BestSn:     utilities.Float(line[11]),
		BestCJ:     utilities.Float(line[12]),
		Total:      utilities.Float(line[13]),
		Sinclair:   0.0,
		Federation: line[14],
	}
	return
}
