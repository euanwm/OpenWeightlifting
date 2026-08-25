package sinclair

import (
	"backend/structs"
	"math"
	"strconv"
)

const (
	naimSinclair = 505 + 1 // The extra 1 is for rounding etc.
	minBW        = 20      // KG, nobody is breaking records at that weight
)

type Year int

type Coefficients struct {
	aMale   float64
	bMale   float64
	aFemale float64
	bFemale float64
}

var CoefficientsByYear = map[Year]Coefficients{
	Year(2021): {
		aMale:   0.722762521,
		bMale:   193.609,
		aFemale: 0.787004341,
		bFemale: 153.757,
	},
	Year(2017): {
		aMale:   0.75194503,
		bMale:   175.508,
		aFemale: 0.783497476,
		bFemale: 153.655,
	},
	Year(2013): {
		aMale:   0.794358141,
		bMale:   174.393,
		aFemale: 0.89726074,
		bFemale: 148.026,
	},
	Year(2009): {
		aMale:   0.784780654,
		bMale:   173.961,
		aFemale: 1.056683941,
		bFemale: 125.441,
	},
	Year(2005): {
		aMale:   0.845716976,
		bMale:   168.091,
		aFemale: 1.316081431,
		bFemale: 107.844,
	},
	Year(2001): {
		aMale:   0.938573813,
		bMale:   135.390,
		aFemale: 1.005487664,
		bFemale: 112.811,
	},
}

func assignCoefficients(entry *structs.Lift) (float64, float64) {

	var entryYear, err = strconv.Atoi(entry.Event.Date[:4])
	var coefficients Coefficients
	if err != nil {
		panic("coefficients year conversion failed")
	}
	switch {
	case entryYear <= 2000:
		coefficients = CoefficientsByYear[Year(2001)]
	case entryYear >= 2001 && entryYear <= 2004:
		coefficients = CoefficientsByYear[Year(2001)]
	case entryYear >= 2005 && entryYear <= 2008:
		coefficients = CoefficientsByYear[Year(2005)]
	case entryYear >= 2009 && entryYear <= 2012:
		coefficients = CoefficientsByYear[Year(2009)]
	case entryYear >= 2013 && entryYear <= 2016:
		coefficients = CoefficientsByYear[Year(2013)]
	case entryYear >= 2017 && entryYear <= 2020:
		coefficients = CoefficientsByYear[Year(2017)]
	case entryYear >= 2021 && entryYear <= 2024:
		coefficients = CoefficientsByYear[Year(2021)]
	case entryYear >= 2025:
		coefficients = CoefficientsByYear[Year(2021)]
	}
	if !entry.Lifter.IsMale() {
		return coefficients.aFemale, coefficients.bFemale
	}
	return coefficients.aMale, coefficients.bMale
}

// CalcSinclair Calculates the sinclair of a result passed to it. We are using ONLY the Senior coefficient because
// the Masters coefficient is absolute nonsense. You'll see there's a lot of switching between float types.
// It's frustrating but it serves a purpose.
func CalcSinclair(result *structs.Lift) {
	var coEffA, coEffB = assignCoefficients(result)

	total := result.Total.Float64()
	bodyweight := result.Bodyweight.Float64()

	// todo: add in error handling
	if bodyweight <= coEffB {
		var X = math.Log10(bodyweight / coEffB)
		var expX = X * X
		var coEffExp = coEffA * expX
		var expSum = math.Pow(10, coEffExp)
		var sinclair = total * expSum
		if sinclair <= naimSinclair {
			result.Sinclair = sinclair
		}
	} else if total <= naimSinclair {
		result.Sinclair = total
	}
}
