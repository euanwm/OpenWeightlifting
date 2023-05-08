package lifter

import (
	"backend/structs"
	"reflect"
	"testing"
)

// todo: add more details to allow more strict testing
var sampleLeaderboardData = &structs.LeaderboardData{
	AllTotals: []structs.Entry{
		{Name: "John Smith", Total: 123},
		{Name: "john smith", Total: 234},
		{Name: "John smoth", Total: 345},
		{Name: "Joanne Smith", Total: 123},
		{Name: "joanne smith", Total: 234},
		{Name: "joanne smith", Total: 235},
		{Name: "joanne Smoth", Total: 345},
	},
}

func TestFetchLifts(t *testing.T) {
	type args struct {
		name        structs.NameSearch
		leaderboard *structs.LeaderboardData
	}
	tests := []struct {
		name             string
		args             args
		wantLifterData   structs.LifterHistory
		expectedLiftsInt int
	}{
		{name: "Single Lift", args: args{
			name:        structs.NameSearch{NameStr: "John Smith"},
			leaderboard: sampleLeaderboardData},
			wantLifterData:   structs.LifterHistory{NameStr: "John Smith"},
			expectedLiftsInt: 1},
		{name: "No Lifts", args: args{
			name:        structs.NameSearch{NameStr: "JOHN SMITH"},
			leaderboard: sampleLeaderboardData},
			wantLifterData:   structs.LifterHistory{NameStr: "JOHN SMITH", Lifts: nil},
			expectedLiftsInt: 0,
		},
		{name: "Multiple Lifts", args: args{
			name:        structs.NameSearch{NameStr: "joanne smith"},
			leaderboard: sampleLeaderboardData},
			wantLifterData:   structs.LifterHistory{NameStr: "joanne smith"},
			expectedLiftsInt: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLifterData := FetchLifts(tt.args.name, tt.args.leaderboard); !reflect.DeepEqual(gotLifterData.NameStr, tt.wantLifterData.NameStr) || !reflect.DeepEqual(len(gotLifterData.Lifts), tt.expectedLiftsInt) {
				t.Errorf("FetchLifts() = %v, want %v", gotLifterData, tt.wantLifterData)
			}
		})
	}
}

func TestNameSearch(t *testing.T) {
	type args struct {
		nameStr  string
		nameList []structs.Entry
	}
	tests := []struct {
		name          string
		args          args
		wantNameSlice []string
	}{
		{name: "Single Match", args: args{
			nameStr: "Dave Smith",
			nameList: []structs.Entry{
				{Name: "andrew smith"},
				{Name: "Dave Smith"},
				{Name: "John Smith"},
				{Name: "maybe john smith"},
			}},
			wantNameSlice: []string{"Dave Smith"},
		},
		{name: "Multiple Match (case insensitive)", args: args{
			nameStr: "John Smith",
			nameList: []structs.Entry{
				{Name: "john smith"},
				{Name: "john not smith"},
				{Name: "John Smith"},
				{Name: "john Smith"},
			}},
			wantNameSlice: []string{"john smith", "John Smith", "john Smith"},
		},
		{name: "No Match on Spelling", args: args{
			nameStr: "John Smith",
			nameList: []structs.Entry{
				{Name: "jim smof"},
				{Name: "dof smith"},
				{Name: "john smof"},
			}},
			wantNameSlice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNamePositions := NameSearch(tt.args.nameStr, &tt.args.nameList); !reflect.DeepEqual(gotNamePositions, tt.wantNameSlice) {
				t.Errorf("NameSearch() = %v, want %v", gotNamePositions, tt.wantNameSlice)
			}
		})
	}
}
