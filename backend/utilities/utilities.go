package utilities

import (
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

// Float - Converts a string containing a float32 to exactly that
func Float(preFloatStr string) (retFloat float32) {
	convFloat, _ := strconv.ParseFloat(preFloatStr, 32)
	retFloat = float32(convFloat)
	return
}

// SliceContains - Returns true if the slice contains the string
func SliceContains(strQuery string, sliceData []string) bool {
	for _, value := range sliceData {
		if value == strQuery {
			return true
		}
	}
	return false
}

// MapContains - Returns true if the map/dict matches the primary/index data
func MapContains(strQuery string, mapData map[string]string) bool {
	for index := range mapData {
		if index == strQuery {
			return true
		}
	}
	return false
}

// Contains - Returns true if a substring within a string exists
func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

// LoadCsvFile Returns the contents of a CSV file as a nested slice minus the header line
func LoadCsvFile(file io.Reader) (csvContents [][]string) {
	reader := csv.NewReader(file)
	csvContents, _ = reader.ReadAll()
	return csvContents[1:]
}

func StringToDate(dateString string) (date time.Time, borkt error) {
	const rfc3339partial string = "T15:04:05Z"
	date, borkt = time.Parse(time.RFC3339, dateString+rfc3339partial)
	if borkt != nil {
		return time.Time{}, borkt
	}
	return date, nil
}

// ReverseSlice - Returns a reversed slice
func ReverseSlice[S any](liftResults []S) (reversedSlice []S) {
	for i := len(liftResults) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, liftResults[i])
	}
	return
}
