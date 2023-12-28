package dbtools

import (
	database "backend/event_data"
	"backend/structs"
	"backend/utilities"
	"io/fs"
	"log"
	"path"
)

func CollateAll(eventmetadata *structs.EventsMetaData) (allData [][]string) {
	dirs := getFedDirs()
	for _, fed := range dirs {
		allResults := loadAllFedEvents(fed, eventmetadata)
		allData = append(allData, allResults...)
	}
	return allData
}

// insertFederation Inserts federation to each event line prior as it's required for the frontend discrimination.
func insertFederation(event [][]string, federation string) [][]string {
	for i := range event {
		event[i] = append(event[i], federation)
	}
	return event
}

// Returns an unsorted nested slice of all events from a single federation/organiser
func loadAllFedEvents(federation string, metadata *structs.EventsMetaData) (allEvents [][]string) {
	allFiles, err := database.Database.ReadDir(federation)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range allFiles {
		func() {
			fileHandle, err := database.Database.Open(path.Join(federation, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			defer func(fileHandle fs.File) {
				err := fileHandle.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(fileHandle)

			eventData := utilities.LoadCsvFile(fileHandle)
			eventData = insertFederation(eventData, federation)
			allEvents = append(allEvents, eventData...)

			metadata.Name = append(metadata.Name, eventData[0][0])
			metadata.Federation = append(metadata.Federation, federation)
			metadata.Date = append(metadata.Date, eventData[0][1])
			metadata.ID = append(metadata.ID, file.Name())
		}()
	}
	return
}

// Returns a slice of the named directories within the database.
// All directories must be named by federation/organiser.
func getFedDirs() (federationDirs []string) {
	dirs, err := database.Database.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		federationDirs = append(federationDirs, dir.Name())
	}
	return federationDirs
}

func LoadSingleEvent(federation, eventID string) (event []structs.Entry) {
	fileHandle, err := database.Database.Open(path.Join(federation, eventID))
	if err != nil {
		log.Println("Error in opening file: ", err)
		return
	}
	defer func(fileHandle fs.File) {
		err := fileHandle.Close()
		if err != nil {
			log.Println("Error in closing file: ", err)
			return
		}
	}(fileHandle)

	eventData := utilities.LoadCsvFile(fileHandle)
	eventData = insertFederation(eventData, federation)
	preEvent, _ := ParseData(eventData)
	event = preEvent.Lifts
	return
}
