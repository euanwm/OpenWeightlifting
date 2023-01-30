package lifter

import "backend/utilities"

var instaHandles = map[string]string{
	"Euan Meston":      "scream_and_jerk",
	"KRYSTAL CAMPBELL": "da.real.krys",
	"shite":            "poo_poo_land",
	"otherShite":       "pee_pee_land",
}

func CheckUserList(lifterName string) string {
	if utilities.MapContains(lifterName, instaHandles) {
		return instaHandles[lifterName]
	}
	return ""
}
