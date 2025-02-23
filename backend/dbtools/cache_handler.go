package dbtools

import (
	"backend/structs"
	"sync"
)

type QueryState int

const (
	None QueryState = iota
	Working
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
	query.Start, query.Stop = 0, 0
	q.HashStore.Store(query, Query{dataPositions, Completed})
}

func (q *QueryCache) InitQuery(query structs.LeaderboardPayload) {
	q.HashStore.Store(query, Query{
		DataPositions: nil,
		Status:        Working,
	})
}

func (q *QueryCache) QueryStatus(query structs.LeaderboardPayload) QueryState {
	query.Start, query.Stop = 0, 0
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
func (q *QueryCache) CheckQuery(query structs.LeaderboardPayload) (state QueryState, positions []int) {
	query.Start, query.Stop = 0, 0
	loadedData, ok := q.HashStore.Load(query)
	if ok {
		storedQuery, ok := loadedData.(Query)
		if ok && (storedQuery.Status == Completed) || (storedQuery.Status == Working) {
			state = storedQuery.Status
			positions = storedQuery.DataPositions
			return
		}
	}

	state = None
	positions = []int{}
	return state, positions
}
