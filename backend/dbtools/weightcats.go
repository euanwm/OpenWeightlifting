package dbtools

import (
	"backend/enum"
	"backend/structs"
)

var WeightClassList = map[string]structs.WeightClass{
	"allcats": {Gender: enum.ALLCATS, Upper: 0, Lower: 0},
	"M55":     {Gender: enum.Male, Upper: 55.00, Lower: 0},
	"M61":     {Gender: enum.Male, Upper: 61.00, Lower: 55.01},
	"M67":     {Gender: enum.Male, Upper: 67.00, Lower: 61.01},
	"M73":     {Gender: enum.Male, Upper: 73.00, Lower: 67.01},
	"M81":     {Gender: enum.Male, Upper: 81.00, Lower: 73.01},
	"M89":     {Gender: enum.Male, Upper: 89.00, Lower: 81.01},
	"M96":     {Gender: enum.Male, Upper: 96.00, Lower: 89.01},
	"M102":    {Gender: enum.Male, Upper: 102.00, Lower: 96.01},
	"M109":    {Gender: enum.Male, Upper: 109.00, Lower: 102.01},
	"M109+":   {Gender: enum.Male, Upper: enum.MaximumBodyweight, Lower: 109.01},
	"F45":     {Gender: enum.Female, Upper: 45.00, Lower: 0},
	"F49":     {Gender: enum.Female, Upper: 49.00, Lower: 45.01},
	"F55":     {Gender: enum.Female, Upper: 55.00, Lower: 49.01},
	"F59":     {Gender: enum.Female, Upper: 59.00, Lower: 55.01},
	"F64":     {Gender: enum.Female, Upper: 64.00, Lower: 59.01},
	"F71":     {Gender: enum.Female, Upper: 71.00, Lower: 64.01},
	"F76":     {Gender: enum.Female, Upper: 76.00, Lower: 71.01},
	"F81":     {Gender: enum.Female, Upper: 81.00, Lower: 76.01},
	"F87":     {Gender: enum.Female, Upper: 87.00, Lower: 81.01},
	"F87+":    {Gender: enum.Female, Upper: enum.MaximumBodyweight, Lower: 87.01},
}
