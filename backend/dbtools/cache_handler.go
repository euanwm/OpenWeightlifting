package dbtools

type QueryCache struct {
	Store []Query
}

type Query struct {
	Filter        StrippedPayload
	DataPositions []int
}

type StrippedPayload struct {
	SortBy      string
	Federation  string
	WeightClass string
	Year        int
	StartDate   string
	EndDate     string
}

// AddQuery - Adds a query to the cache.
func (q *QueryCache) AddQuery(query StrippedPayload, dataPositions []int) {
	q.Store = append(q.Store, Query{Filter: query, DataPositions: dataPositions})
}

// CheckQuery - Checks if the query has been run before, if so, return the data positions.
func (q *QueryCache) CheckQuery(query StrippedPayload) (bool, []int) {
	for _, c := range q.Store {
		if c.Filter == query {
			return true, c.DataPositions
		}
	}
	return false, nil
}
