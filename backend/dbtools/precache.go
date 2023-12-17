package dbtools

import (
	"backend/enum"
	"backend/structs"
)

// PreCacheQuery - Queries to pre-cache on startup.
var PreCacheQuery = []structs.LeaderboardPayload{
	{
		SortBy:      "total",
		Federation:  "allfeds",
		WeightClass: "MALL",
		StartDate:   enum.ZeroDate,
		EndDate:     enum.MaxDate,
	},
	{
		SortBy:      "total",
		Federation:  "allfeds",
		WeightClass: "FALL",
		StartDate:   enum.ZeroDate,
		EndDate:     enum.MaxDate,
	},
	{
		SortBy:      "sinclair",
		Federation:  "allfeds",
		WeightClass: "MALL",
		StartDate:   enum.ZeroDate,
		EndDate:     enum.MaxDate,
	},
	{
		SortBy:      "sinclair",
		Federation:  "allfeds",
		WeightClass: "FALL",
		StartDate:   enum.ZeroDate,
		EndDate:     enum.MaxDate,
	},
}
