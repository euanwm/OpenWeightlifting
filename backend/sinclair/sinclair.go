package sinclair

import (
	"backend/structs"
	"math"
)

// Coefficient numbers
const (
	aMale   = 0.751945030
	bMale   = 175.508
	aFemale = 0.783497476
	bFemale = 153.655
)

// CalcSinclair Calculates the sinclair of a result passed to it. We are using ONLY the Senior coefficient because
// the Masters coefficient is absolute nonsense. You'll see there's a lot of switching between float types.
// It's frustrating but it serves a purpose.
func CalcSinclair(result *structs.Entry, male bool) {
	var coEffA = aMale
	var coEffB = bMale
	if male == false {
		coEffA = aFemale
		coEffB = bFemale
	}
	if result.Total != 0 && result.Bodyweight >= 0 {
		if float64(result.Bodyweight) <= coEffB {
			var X = math.Log10(float64(result.Bodyweight) / coEffB)
			var expX = math.Pow(X, 2)
			var coEffExp = coEffA * expX
			var expSum = math.Pow(10, coEffExp)
			result.Sinclair = float32(float64(result.Total) * expSum)
		} else {
			result.Sinclair = result.Total
		}
	}
}
