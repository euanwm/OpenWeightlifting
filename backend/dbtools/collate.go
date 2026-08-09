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

// columnAliases maps each canonical column position (consumed by assignStruct) to
// one or more possible header names found across different federation CSV formats.
// The federation column is appended separately by insertFederation.
var columnAliases = [][]string{
	{"event", "meet"},                      // 0: Event
	{"date"},                               // 1: Date
	{"gender", "age_category", "category"}, // 2: Gender
	{"lifter", "lifter_name"},              // 3: Name
	{"body_weight_(kg)", "bodyweight"},     // 4: Bodyweight
	{"snatch_lift_1", "snatch_1"},          // 5: Sn1
	{"snatch_lift_2", "snatch_2"},          // 6: Sn2
	{"snatch_lift_3", "snatch_3"},          // 7: Sn3
	{"c&j_lift_1", "cj_1"},                 // 8: CJ1
	{"c&j_lift_2", "cj_2"},                 // 9: CJ2
	{"c&j_lift_3", "cj_3"},                 // 10: CJ3
	{"best_snatch"},                        // 11: BestSn
	{"best_c&j", "best_cj"},                // 12: BestCJ
	{"total"},                              // 13: Total
}

// normalizeColumns reorders rows to match columnAliases order using the CSV header,
// allowing CSVs with additional, renamed, or differently-ordered columns to be handled gracefully.
// Columns not matched by any alias default to an empty string.
func normalizeColumns(header []string, rows [][]string) [][]string {
	colIndex := make(map[string]int, len(header))
	for i, h := range header {
		colIndex[h] = i
	}
	// Resolve each canonical position to its source column index, using the first matching alias.
	resolved := make([]int, len(columnAliases))
	for j, aliases := range columnAliases {
		resolved[j] = -1
		for _, alias := range aliases {
			if idx, ok := colIndex[alias]; ok {
				resolved[j] = idx
				break
			}
		}
	}
	normalized := make([][]string, len(rows))
	for i, row := range rows {
		normalizedRow := make([]string, len(columnAliases))
		for j, srcIdx := range resolved {
			if srcIdx >= 0 && srcIdx < len(row) {
				normalizedRow[j] = row[srcIdx]
			}
		}
		normalized[i] = normalizedRow
	}
	return normalized
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

			header, eventData := utilities.LoadCsvFile(fileHandle)
			if len(eventData) == 0 {
				return
			}
			eventData = normalizeColumns(header, eventData)
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

func createSingleEvent(federation, filename string, eventsData *structs.EventsData, allLifts *structs.AllLifts, lifterRoster *structs.LifterRoster) {
	fileHandle, err := database.Database.Open(path.Join(federation, filename))
	if err != nil {
		log.Fatal(err)
	}
	defer func(fileHandle fs.File) {
		err := fileHandle.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fileHandle)

	header, eventData := utilities.LoadCsvFile(fileHandle)
	if len(eventData) == 0 {
		return
	}
	eventData = normalizeColumns(header, eventData)
	event := &structs.Event{
		Federation: federation,
		CSVID:      filename,
	}

	// process all lifts
	liftPtrSlice := make([]*structs.Lift, 0, len(eventData))
	for _, row := range eventData {
		lift := &structs.Lift{
			Bodyweight: structs.NewWeightKgFromString(row[4]),
			Sn1:        structs.NewWeightKgFromString(row[5]),
			Sn2:        structs.NewWeightKgFromString(row[6]),
			Sn3:        structs.NewWeightKgFromString(row[7]),
			CJ1:        structs.NewWeightKgFromString(row[8]),
			CJ2:        structs.NewWeightKgFromString(row[9]),
			CJ3:        structs.NewWeightKgFromString(row[10]),
			BestSn:     structs.NewWeightKgFromString(row[11]),
			BestCJ:     structs.NewWeightKgFromString(row[12]),
			Total:      structs.NewWeightKgFromString(row[13]),
			Sinclair:   0.0,
			Category:   row[2],
			Event:      event,
		}
		lift.Lifter = lifterRoster.Add(row[3], row[2], federation, lift)
		event.Name = row[0]
		event.Date = row[1]
		liftPtrSlice = append(liftPtrSlice, lift)
	}
	event.Results = liftPtrSlice
	allLifts.Lifts = append(allLifts.Lifts, liftPtrSlice...)
	eventsData.Events = append(eventsData.Events, event)
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

	header, eventData := utilities.LoadCsvFile(fileHandle)
	eventData = normalizeColumns(header, eventData)
	eventData = insertFederation(eventData, federation)
	preEvent, _ := ParseData(eventData)
	event = preEvent.Lifts
	return
}
