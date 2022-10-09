package dbtools

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

//Float - Converts a string containing a float32 to exactly that
func Float(preFloatStr string) (retFloat float32) {
	convFloat, _ := strconv.ParseFloat(preFloatStr, 32)
	retFloat = float32(convFloat)
	return
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

//WriteCSV Writes CSV file, first arg is the filepath/name. Second is the bigSlice data.
func WriteCSV(csvFp string, bigSlice [][]string) {
	newCsvFile, err := os.Create(csvFp)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(newCsvFile)
	writeData := writer.WriteAll(bigSlice)
	if writeData != nil {
		fmt.Println(writeData)
	}
}
