package dbtools

import (
	"backend/structs"
	"testing"
)

func TestQueryCache_AddQuery(t *testing.T) {
	type fields struct {
		Store []Query
	}
	type args struct {
		query         structs.LeaderboardPayload
		dataPositions []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"AddQuery", fields{Store: []Query{}}, args{query: structs.LeaderboardPayload{}, dataPositions: []int{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QueryCache{
				Store: tt.fields.Store,
			}
			q.AddQuery(tt.args.query, tt.args.dataPositions)
		})
	}
}

func TestQueryCache_CheckQuery(t *testing.T) {
	type fields struct {
		Store []Query
	}
	type args struct {
		query structs.LeaderboardPayload
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  []int
	}{
		{"CheckQuery", fields{Store: []Query{{Filter: structs.LeaderboardPayload{SortBy: "total", Federation: "IPF", WeightClass: "93", StartDate: "2018-01-01", EndDate: "2018-01-01"}, DataPositions: []int{1}}}}, args{query: structs.LeaderboardPayload{SortBy: "total", Federation: "IPF", WeightClass: "93", StartDate: "2018-01-01", EndDate: "2018-01-01"}}, true, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &QueryCache{
				Store: tt.fields.Store,
			}
			got, got1 := q.CheckQuery(tt.args.query)
			if got != tt.want {
				t.Errorf("QueryCache.CheckQuery() got = %v, want %v", got, tt.want)
			}
			if len(got1) != len(tt.want1) {
				t.Errorf("QueryCache.CheckQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
