package lifter

import "backend/utilities"

var instaHandles = map[string]string{
	"Euan Meston":      "scream_and_jerk",
	"KRYSTAL CAMPBELL": "da.real.krys",
	"TALAKHADZE Lasha": "talakhadzelasha_official",
}

func CheckUserList(lifterName string) (bool, string) {
	if utilities.MapContains(lifterName, instaHandles) {
		return true, instaHandles[lifterName]
	}
	return false, ""
}
