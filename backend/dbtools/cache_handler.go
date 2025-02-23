package dbtools

import (
	"backend/enum"
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

// CheckQuery - Checks if the query has been run before, if so, return the query state and data positions if they exist
func (q *QueryCache) CheckQuery(query structs.LeaderboardPayload) (QueryState, []int) {
	loadedData, ok := q.HashStore.Load(query)
	if ok {
		storedQuery, ok := loadedData.(Query)
		if ok {
			return storedQuery.Status, storedQuery.DataPositions
		}
	}

	state := None
	var positions []int
	q.HashStore.Range(func(key, value interface{}) bool {
		loadedQuery, _ := value.(Query)
		queryPayload, _ := key.(structs.LeaderboardPayload)
		if queryPayload.SortBy == query.SortBy && queryPayload.Federation == query.Federation && queryPayload.WeightClass == query.WeightClass && queryPayload.Year == enum.AllYearsStr {
			positions = loadedQuery.DataPositions
			state = Partial
			return false
		}
		if queryPayload.SortBy == query.SortBy && queryPayload.Federation == query.Federation && queryPayload.Year == enum.AllYearsStr {
			if query.WeightClass[0] == 'M' && queryPayload.WeightClass == "MALL" {
				positions = loadedQuery.DataPositions
				state = Partial
				return false
			}
			if query.WeightClass[0] == 'F' && queryPayload.WeightClass == "FALL" {
				positions = loadedQuery.DataPositions
				state = Partial
				return false
			}
		}
		return false
	})
	return state, positions
}
