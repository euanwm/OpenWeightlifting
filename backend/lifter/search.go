package lifter

import "strings"

// NameSearch takes a partial string and returns a slice of positions within the AllNames slice that could be a match
func NameSearch(nameStr string, nameList *[]string) (namePositions []int) {
	nameStr = strings.ToLower(nameStr)
	for pos, name := range *nameList {
		if strings.Contains(strings.ToLower(name), nameStr) {
			namePositions = append(namePositions, pos)
		}
	}
	return
}
