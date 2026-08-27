package dbtools

import (
	"backend/enum"
	"backend/structs"
)

type RulesChecker struct {
	MaxSnatch       structs.FixedFloat
	MaxCleanAndJerk structs.FixedFloat
	MaxTotal        structs.FixedFloat
	MaxBodyweight   structs.FixedFloat
	MinBodyweight   structs.FixedFloat
}

func (e *RulesChecker) PassesRules(lift *structs.Lift) bool {
	if !lift.Total.GreaterThanOrEqual(structs.NewFixedFloat(0)) {
		return false
	}
	return lift.Total.LessThan(e.MaxTotal) && lift.Bodyweight.GreaterThan(e.MinBodyweight)
}

func LoadRules() *RulesChecker {
	return &RulesChecker{
		MaxSnatch:       structs.NewFixedFloat(enum.MaxSnatch),
		MaxCleanAndJerk: structs.NewFixedFloat(enum.MaxCleanAndJerk),
		MaxTotal:        structs.NewFixedFloat(enum.MaxTotal),
		MaxBodyweight:   structs.NewFixedFloat(enum.MaximumBodyweight),
		MinBodyweight:   structs.NewFixedFloat(enum.MinimumBodyweight),
	}
}
