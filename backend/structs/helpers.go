package structs

import "backend/enum"

func IterateFloatSlice(data []Entry, item string) (floatSl []float32) {
	// todo: implement DRY principle
	switch item {
	case enum.Total:
		for _, lift := range data {
			floatSl = append(floatSl, lift.Total.Float32())
		}
	case enum.BestSnatch:
		for _, lift := range data {
			floatSl = append(floatSl, lift.BestSn.Float32())
		}
	case enum.BestCJ:
		for _, lift := range data {
			floatSl = append(floatSl, lift.BestCJ.Float32())
		}
	case enum.Bodyweight:
		for _, lift := range data {
			floatSl = append(floatSl, lift.Bodyweight.Float32())
		}
	}
	return
}
