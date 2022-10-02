package dbtools

import (
	"backend/enum"
	"backend/sinclair"
	"backend/structs"
	"regexp"
	"strconv"
)

// SortGender Splits results into 3 categories, male, female, and unknown.
func SortGender(bigData [][]string) (male []structs.Entry, female []structs.Entry, unknown []structs.Entry) {
	for _, contents := range bigData {
		dataStruct := assignStruct(contents)
		gender := setGender(&dataStruct)
		switch gender {
		case enum.Male:
			if dataStruct.Total > 0 && dataStruct.Bodyweight > 0 {
				sinclair.CalcSinclair(&dataStruct, true)
			}
			male = append(male, dataStruct)
		case enum.Female:
			if dataStruct.Total > 0 && dataStruct.Bodyweight > 0 {
				sinclair.CalcSinclair(&dataStruct, false)
			}
			female = append(female, dataStruct)
		case enum.Unknown:
			unknown = append(unknown, dataStruct)
		}
	}
	return
}

func setGender(entry *structs.Entry) (gender string) {
	if entry.Gender == enum.Male || regGenderCheck(entry) == enum.Male {
		return enum.Male
	} else if entry.Gender == enum.Female || regGenderCheck(entry) == enum.Female {
		return enum.Female
	} else {
		return enum.Unknown
	}
}

func regGenderCheck(entry *structs.Entry) (gender string) {
	searchMale, _ := regexp.Match("Men", []byte(entry.Gender))
	searchFemale, _ := regexp.Match("Women", []byte(entry.Gender))
	if searchMale == true {
		return enum.Male
	} else if searchFemale == true {
		return enum.Female
	} else {
		return enum.Unknown
	}
}

func assignStruct(line []string) (lineStruct structs.Entry) {
	floatTotal, _ := strconv.ParseFloat(line[13], 32)
	floatBodyweight, _ := strconv.ParseFloat(line[4], 32)
	lineStruct = structs.Entry{
		Event:      line[0],
		Date:       line[1],
		Gender:     line[2],
		Name:       line[3],
		Bodyweight: float32(floatBodyweight),
		Sn1:        line[5],
		Sn2:        line[6],
		Sn3:        line[7],
		CJ1:        line[8],
		CJ2:        line[9],
		CJ3:        line[10],
		BestSn:     line[11],
		BestCJ:     line[12],
		Total:      float32(floatTotal),
		Sinclair:   0.0,
		Federation: line[14],
	}
	return
}
