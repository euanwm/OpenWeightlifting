package dbtools

import (
	"backend/enum"
	"backend/structs"
)

type RulesChecker struct {
	MaxSnatch       structs.WeightKg
	MaxCleanAndJerk structs.WeightKg
	MaxTotal        structs.WeightKg
	MaxBodyweight   structs.WeightKg
	MinBodyweight   structs.WeightKg
}

func (e *RulesChecker) PassesRules(lift *structs.Lift) bool {
	if !lift.Total.GreaterThanOrEqual(structs.NewWeightKg(0)) {
		return false
	}
	return lift.Total.LessThan(e.MaxTotal) && lift.Bodyweight.GreaterThan(e.MinBodyweight)
}

func LoadRules() *RulesChecker {
	return &RulesChecker{
		MaxSnatch:       structs.NewWeightKg(enum.MaxSnatch),
		MaxCleanAndJerk: structs.NewWeightKg(enum.MaxCleanAndJerk),
		MaxTotal:        structs.NewWeightKg(enum.MaxTotal),
		MaxBodyweight:   structs.NewWeightKg(enum.MaximumBodyweight),
		MinBodyweight:   structs.NewWeightKg(enum.MinimumBodyweight),
	}
}
