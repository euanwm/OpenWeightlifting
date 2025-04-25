package dbtools

import (
	"backend/enum"
	"backend/sinclair"
	"backend/structs"
	"log"
	"strings"
)

// ParseData Splits results into 3 categories, male, female, and unknown.
func ParseData(bigData [][]string) (allLifts structs.AllData, unknown structs.AllData) {
	max_total := structs.NewWeightKg(float64(enum.MaxTotal))
	min_bodyweight := structs.NewWeightKg(float64(enum.MinimumBodyweight))

	for _, contents := range bigData {
		dataStruct, valid := assignStruct(contents)
		if !valid {
			continue
		}

		gender := getGender(&dataStruct)
		switch gender {
		case enum.Male:
			if dataStruct.Total.IsPositive() &&
				dataStruct.Total.LessThan(max_total) &&
				dataStruct.Bodyweight.GreaterThan(min_bodyweight) {
				// todo: add in error handling for CalcSinclair
				sinclair.CalcSinclair(&dataStruct, true)
			}
			allLifts.Lifts = append(allLifts.Lifts, dataStruct)
		case enum.Female:
			if dataStruct.Total.IsPositive() &&
				dataStruct.Total.LessThan(max_total) &&
				dataStruct.Bodyweight.GreaterThan(min_bodyweight) {
				sinclair.CalcSinclair(&dataStruct, false)
			}
			allLifts.Lifts = append(allLifts.Lifts, dataStruct)
		case enum.Unknown:
			unknown.Lifts = append(unknown.Lifts, dataStruct)
		}
	}
	return
}

func getGender(entry *structs.Entry) (gender string) {
	switch {
	case entry.Gender == enum.Male:
		return enum.Male
	case entry.Gender == enum.Female:
		return enum.Female
	case strings.Contains(entry.Gender, "Men"):
		return enum.Male
	case strings.Contains(entry.Gender, "Women"):
		return enum.Female
	default:
		return enum.Unknown
	}
}

func assignStruct(line []string) (lineStruct structs.Entry, valid bool) {
	if line[0][0] == '#' {
		log.Print("Skipping entry: ", line)
		return lineStruct, false
	}
	lineStruct = structs.Entry{
		Event:      line[0],
		Date:       line[1],
		Gender:     line[2],
		Name:       line[3],
		Bodyweight: structs.NewWeightKgFromString(line[4]),
		Sn1:        structs.NewWeightKgFromString(line[5]),
		Sn2:        structs.NewWeightKgFromString(line[6]),
		Sn3:        structs.NewWeightKgFromString(line[7]),
		CJ1:        structs.NewWeightKgFromString(line[8]),
		CJ2:        structs.NewWeightKgFromString(line[9]),
		CJ3:        structs.NewWeightKgFromString(line[10]),
		BestSn:     structs.NewWeightKgFromString(line[11]),
		BestCJ:     structs.NewWeightKgFromString(line[12]),
		Total:      structs.NewWeightKgFromString(line[13]),
		Sinclair:   0.0,
		Federation: line[14],
	}
	return lineStruct, true
}
