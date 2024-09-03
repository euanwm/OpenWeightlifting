package sinclair

import (
	"backend/structs"
	"math"
)

// Coefficient numbers
const (
	// These are the 2021-2024 Coefficient numbers
	aMale        = 0.722762521
	bMale        = 193.609
	aFemale      = 0.787004341
	bFemale      = 153.757
	naimSinclair = 505 + 1 // The extra 1 is for rounding etc.
	minBW        = 20      // KG, nobody is breaking records at that weight
)

// CalcSinclair Calculates the sinclair of a result passed to it. We are using ONLY the Senior coefficient because
// the Masters coefficient is absolute nonsense. You'll see there's a lot of switching between float types.
// It's frustrating but it serves a purpose.
func CalcSinclair(result *structs.Entry, male bool) {
	// Fast path: a zero or negative total has a zero Sinclair.
	if !result.Total.IsPositive() {
		result.Sinclair = 0
		return
	}

	// A bodyweight below the cutoff also receives a zero score.
	if result.Bodyweight.LessThanOrEqual(structs.NewWeightKgFromInt32(minBW)) {
		result.Sinclair = 0
		return
	}

	var coEffA = aMale
	var coEffB = bMale
	if !male {
		coEffA = aFemale
		coEffB = bFemale
	}

	total := result.Total.Float64()
	bodyweight := result.Bodyweight.Float64()

	// todo: add in error handling
	if bodyweight <= coEffB {
		var X = math.Log10(bodyweight / coEffB)
		var expX = math.Pow(X, 2)
		var coEffExp = coEffA * expX
		var expSum = math.Pow(10, coEffExp)
		var sinclair = float32(total * expSum)
		if sinclair <= naimSinclair {
			result.Sinclair = sinclair
		}
	} else if total <= naimSinclair {
		result.Sinclair = float32(total)
	}
}
