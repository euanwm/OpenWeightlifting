package dbtools

import (
	"backend/structs"
	"sync"
)

type QueryState int

const (
	None QueryState = iota
	Working
	Partial
	Completed
)

type QueryCache struct {
	HashStore sync.Map // [structs.LeaderboardPayload]Query
}

type Query struct {
	DataPositions []int
	Status        QueryState
}

// AddQuery - Adds a query to the cache.
func (q *QueryCache) AddQuery(query structs.LeaderboardPayload, dataPositions []int) {
	q.HashStore.Store(query, Query{dataPositions, Completed})
}

func (q *QueryCache) InitQuery(query structs.LeaderboardPayload) {
	q.HashStore.Store(query, Query{
		DataPositions: nil,
		Status:        Working,
	})
}

func (q *QueryCache) QueryStatus(query structs.LeaderboardPayload) QueryState {
	queryStuff, ok := q.HashStore.Load(query)
	if !ok {
		return None
	} else {
		query, ok := queryStuff.(Query)
		if !ok {
			panic("how the fuck did you fuck this up?")
		}
		return query.Status
	}
}

// CheckQuery - Checks if the query has been run before, if so, return the data positions.
// A true value indicates that the query has been run before.
// A false value indicates that the query has not been run before.
// And a false value with a non-nil slice indicates that a similar query has been run before and the data positions are returned that should be used to filer from.
// Hoping that the last case makes things a bit faster to reduce having to run multiple startup caching queries.
func (q *QueryCache) CheckQuery(query structs.LeaderboardPayload) (QueryState, []int) {
	loadedData, ok := q.HashStore.Load(query)
	if ok {
		storedQuery, ok := loadedData.(Query)
		if ok {
			return storedQuery.Status, storedQuery.DataPositions
		}
	}
	/*
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
	*/
	return None, nil
}
