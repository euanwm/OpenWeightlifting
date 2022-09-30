package dbtools

import (
	"log"
	"os"
	"path"
)

const databaseRoot string = "database_root"

func CollateAll() (allData [][]string) {
	dirs := getFedDirs()
	for _, fed := range dirs {
		var allResults [][]string
		allResults = loadAllFedEvents(fed)
		allData = append(allData, allResults...)
	}
	return allData
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
		eventData := LoadCsvFile(federationPath, file.Name(), true)
		eventData = insertFederation(eventData, federation)
		allEvents = append(allEvents, eventData...)
	}
	return
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
