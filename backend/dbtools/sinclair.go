package dbtools

import (
	"math"
)

// Coefficient numbers
const (
	aMale   = 0.751945030
	bMale   = 175.508
	aFemale = 0.783497476
	bFemale = 153.655
)

// CalcSinclair Calculates the sinclair of a result passed to it. todo: add in the struct shit bit
func CalcSinclair(total float64, bodyweight float64, male bool) (sinclairScore float64) {
	var coEffA = aMale
	var coEffB = bMale
	if male == false {
		coEffA = aFemale
		coEffB = bFemale
	}
	if bodyweight <= coEffB {
		var X = math.Log10(bodyweight / coEffB)
		var expX = math.Pow(X, 2)
		var coEffExp = coEffA * expX
		var expSum = math.Pow(10, coEffExp)
		sinclairScore = total * expSum
	} else {
		sinclairScore = total
	}
	return sinclairScore
}
