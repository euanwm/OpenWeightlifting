package dbtools

import (
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
func (q *QueryCache) CheckQuery(query structs.LeaderboardPayload) (bool, []int) {
	for _, cacheQuery := range q.Store {
		if cacheQuery.Filter.SortBy == query.SortBy && cacheQuery.Filter.Federation == query.Federation && cacheQuery.Filter.WeightClass == query.WeightClass && cacheQuery.Filter.StartDate == query.StartDate && cacheQuery.Filter.EndDate == query.EndDate {
			return true, cacheQuery.DataPositions
		}
	}
	return false, nil
}
