package lifter

import (
	igDatabase "backend/lifterdata"
	"backend/utilities"
	"io/fs"
	"log"
)

func CheckUserList(lifterName string, lifterProfiles map[string]string) (bool, string) {
	if utilities.MapContains(lifterName, lifterProfiles) {
		return true, lifterProfiles[lifterName]
	}
	return false, ""
}

func Build() *map[string]string {
	var lifterGrams [][]string
	var lifterGramsMap = make(map[string]string)
	func() {
		fileHandle, err := igDatabase.InstagramDatabase.Open("ighandles.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer func(fileHandle fs.File) {
			err := fileHandle.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(fileHandle)
		gramData := utilities.LoadCsvFile(fileHandle)
		lifterGrams = append(lifterGrams, gramData...)
	}()

	for _, gram := range lifterGrams {
		lifterGramsMap[gram[0]] = gram[1]
	}
	return &lifterGramsMap
}
