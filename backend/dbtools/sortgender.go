package dbtools

import (
	"regexp"
	"strconv"
)

const (
	Male     string = "male"
	Female   string = "female"
	Unknown  string = "unknown"
	Total    string = "total"
	Sinclair string = "sinclair"
)

// SortGender Splits results into 3 categories, male, female, and unknown then sorts the male and female by descending total
func SortGender(bigData [][]string) (male []Entry, female []Entry, unknown []Entry) {
	for _, contents := range bigData {
		dataStruct := assignStruct(contents)
		gender := setGender(&dataStruct)
		switch gender {
		case Male:
			male = append(male, dataStruct)
		case Female:
			female = append(female, dataStruct)
		case Unknown:
			unknown = append(unknown, dataStruct)
		}
	}
	return
}

func setGender(entry *Entry) (gender string) {
	if entry.Gender == Male || regGenderCheck(entry) == Male {
		return Male
	} else if entry.Gender == Female || regGenderCheck(entry) == Female {
		return Female
	} else {
		return Unknown
	}
}

func regGenderCheck(entry *Entry) (gender string) {
	searchMale, _ := regexp.Match("Men", []byte(entry.Gender))
	searchFemale, _ := regexp.Match("Women", []byte(entry.Gender))
	if searchMale == true {
		return Male
	} else if searchFemale == true {
		return Female
	} else {
		return Unknown
	}
}

func assignStruct(line []string) (lineStruct Entry) {
	floatTotal, _ := strconv.ParseFloat(line[13], 32)
	floatBodyweight, _ := strconv.ParseFloat(line[4], 32)
	lineStruct = Entry{
		Event:      line[0],
		Date:       line[1],
		Gender:     line[2],
		Name:       line[3],
		Bodyweight: floatBodyweight,
		Sn1:        line[5],
		Sn2:        line[6],
		Sn3:        line[7],
		CJ1:        line[8],
		CJ2:        line[9],
		CJ3:        line[10],
		BestSn:     line[11],
		BestCJ:     line[12],
		Total:      floatTotal,
		Sinclair:   0.0,
		Federation: line[14],
	}
	return
}
