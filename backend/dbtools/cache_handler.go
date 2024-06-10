package dbtools

import (
	"backend/enum"
	"backend/structs"
)

type QueryCache struct {
	Store []Query
}

type Query struct {
	Filter        structs.LeaderboardPayload
	DataPositions []int
}

// AddQuery - Adds a query to the cache.
func (q *QueryCache) AddQuery(query structs.LeaderboardPayload, dataPositions []int) {
	q.Store = append(q.Store, Query{Filter: query, DataPositions: dataPositions})
}

// CheckQuery - Checks if the query has been run before, if so, return the data positions.
// A true value indicates that the query has been run before.
// A false value indicates that the query has not been run before.
// And a false value with a non-nil slice indicates that a similar query has been run before and the data positions are returned that should be used to filer from.
// Hoping that the last case makes things a bit faster to reduce having to run multiple startup caching queries.
func (q *QueryCache) CheckQuery(query structs.LeaderboardPayload) (bool, []int) {
	for _, cacheQuery := range q.Store {
		if cacheQuery.Filter == query {
			return true, cacheQuery.DataPositions
		}
		// checks for exact match
		if cacheQuery.Filter.SortBy == query.SortBy && cacheQuery.Filter.Federation == query.Federation && cacheQuery.Filter.WeightClass == query.WeightClass && cacheQuery.Filter.Year == query.Year {
			return true, cacheQuery.DataPositions
		}
	}
	// if we get here, we haven't found a match, so we'll do some partial matching
	for _, cacheQuery := range q.Store {
		// all years for the same total/sinclair, federation and weight class
		if cacheQuery.Filter.SortBy == query.SortBy && cacheQuery.Filter.Federation == query.Federation && cacheQuery.Filter.WeightClass == query.WeightClass && cacheQuery.Filter.Year == enum.AllYearsStr {
			return false, cacheQuery.DataPositions
		}
		// all years for the same total/sinclair, federation, and all gendered weight classes
		if cacheQuery.Filter.SortBy == query.SortBy && cacheQuery.Filter.Federation == query.Federation && cacheQuery.Filter.Year == enum.AllYearsStr {
			if query.WeightClass[0] == 'M' && cacheQuery.Filter.WeightClass == "MALL" {
				return false, cacheQuery.DataPositions
			}
			if query.WeightClass[0] == 'F' && cacheQuery.Filter.WeightClass == "FALL" {
				return false, cacheQuery.DataPositions
			}
		}
	}
	return false, nil
}
