package dbtools

import (
	"encoding/csv"
	"log"
	"os"
	"path"
)

// LoadCsvFile Returns the contents of a CSV file as a nested slice with an option to skip the header line but in a lazy AF way
func LoadCsvFile(folder string, filename string, skipHeader bool) (csvContents [][]string) {
	filepath := path.Join(folder, filename)
	openFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(openFile)
	csvContents, _ = reader.ReadAll()
	if skipHeader == true {
		return csvContents[1:]
	}
	return csvContents
}
