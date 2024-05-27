package dbtools

import (
	"backend/enum"
	"backend/structs"
)

var WeightClassList = map[string]structs.WeightClass{
	"MALL":  {Gender: enum.Male, Upper: structs.NewWeightKg(float64(enum.MaximumBodyweight)), Lower: structs.NewWeightKg(0)},
	"M55":   {Gender: enum.Male, Upper: structs.NewWeightKg(55.00), Lower: structs.NewWeightKg(0)},
	"M61":   {Gender: enum.Male, Upper: structs.NewWeightKg(61.00), Lower: structs.NewWeightKg(55.01)},
	"M67":   {Gender: enum.Male, Upper: structs.NewWeightKg(67.00), Lower: structs.NewWeightKg(61.01)},
	"M73":   {Gender: enum.Male, Upper: structs.NewWeightKg(73.00), Lower: structs.NewWeightKg(67.01)},
	"M81":   {Gender: enum.Male, Upper: structs.NewWeightKg(81.00), Lower: structs.NewWeightKg(73.01)},
	"M89":   {Gender: enum.Male, Upper: structs.NewWeightKg(89.00), Lower: structs.NewWeightKg(81.01)},
	"M96":   {Gender: enum.Male, Upper: structs.NewWeightKg(96.00), Lower: structs.NewWeightKg(89.01)},
	"M102":  {Gender: enum.Male, Upper: structs.NewWeightKg(102.00), Lower: structs.NewWeightKg(96.01)},
	"M109":  {Gender: enum.Male, Upper: structs.NewWeightKg(109.00), Lower: structs.NewWeightKg(102.01)},
	"M109+": {Gender: enum.Male, Upper: structs.NewWeightKg(float64(enum.MaximumBodyweight)), Lower: structs.NewWeightKg(109.01)},
	"FALL":  {Gender: enum.Female, Upper: structs.NewWeightKg(float64(enum.MaximumBodyweight)), Lower: structs.NewWeightKg(0)},
	"F45":   {Gender: enum.Female, Upper: structs.NewWeightKg(45.00), Lower: structs.NewWeightKg(0)},
	"F49":   {Gender: enum.Female, Upper: structs.NewWeightKg(49.00), Lower: structs.NewWeightKg(45.01)},
	"F55":   {Gender: enum.Female, Upper: structs.NewWeightKg(55.00), Lower: structs.NewWeightKg(49.01)},
	"F59":   {Gender: enum.Female, Upper: structs.NewWeightKg(59.00), Lower: structs.NewWeightKg(55.01)},
	"F64":   {Gender: enum.Female, Upper: structs.NewWeightKg(64.00), Lower: structs.NewWeightKg(59.01)},
	"F71":   {Gender: enum.Female, Upper: structs.NewWeightKg(71.00), Lower: structs.NewWeightKg(64.01)},
	"F76":   {Gender: enum.Female, Upper: structs.NewWeightKg(76.00), Lower: structs.NewWeightKg(71.01)},
	"F81":   {Gender: enum.Female, Upper: structs.NewWeightKg(81.00), Lower: structs.NewWeightKg(76.01)},
	"F87":   {Gender: enum.Female, Upper: structs.NewWeightKg(87.00), Lower: structs.NewWeightKg(81.01)},
	"F87+":  {Gender: enum.Female, Upper: structs.NewWeightKg(float64(enum.MaximumBodyweight)), Lower: structs.NewWeightKg(87.01)},
}
