package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
)

const databaseRoot string = "database_root"

func main() {
	dirs := getFedDirs()
	var allData [][]string
	for _, fed := range dirs {
		var allResults [][]string
		allResults = loadAllFedEvents(fed)
		allData = append(allData, allResults...)
	}
	writeCSV("alldata.csv", allData)
}

//Inserts federation to each event line prior as it's required for the frontend discrimination.
func insertFederation(event [][]string, federation string) (newEventData [][]string) {
	for _, line := range event {
		line = append(line, federation)
		newEventData = append(newEventData, line)
	}
	return
}

//Returns an unsorted nested slice of all events from a single federation/organiser
func loadAllFedEvents(federation string) (allEvents [][]string) {
	federationPath := path.Join(databaseRoot, federation)
	allFiles, err := os.ReadDir(federationPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range allFiles {
		eventData := loadCsvFile(federationPath, file.Name(), true)
		eventData = insertFederation(eventData, federation)
		allEvents = append(allEvents, eventData...)
	}
	return
}

//Returns the contents of a CSV file as a nested slice with an option to skip the header line but in a lazy AF way
func loadCsvFile(federation string, filename string, skipHeader bool) (csvContents [][]string) {
	filepath := path.Join(federation, filename)
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

//writeCSV Writes CSV file, first arg is the filepath/name. Second is the bigSlice data.
func writeCSV(csvFp string, bigSlice [][]string) {
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

//Returns a slice of the named directories within the database.
//All directories must be named by federation/organiser.
func getFedDirs() (federationDirs []string) {
	dirs, err := os.ReadDir(databaseRoot)
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		federationDirs = append(federationDirs, dir.Name())
	}
	return federationDirs
}
