package utilities

import (
	"encoding/csv"
	"io"
	"strconv"
)

//Float - Converts a string containing a float32 to exactly that
func Float(preFloatStr string) (retFloat float32) {
	convFloat, _ := strconv.ParseFloat(preFloatStr, 32)
	retFloat = float32(convFloat)
	return
}

//MapContains - Returns true if the map/dict matches the primary/index data
func MapContains(StrQuery string, mapData map[string]string) bool {
	for index, _ := range mapData {
		if index == StrQuery {
			return true
		}
	}
	return false
}

//Contains - Returns true if a substring within a string exists
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
