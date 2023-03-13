package lifter

import (
	"backend/structs"
	"reflect"
	"testing"
)

// todo: add more details to allow more strict testing
var sampleLeaderboardData = &structs.LeaderboardData{
	MaleTotals: []structs.Entry{
		{Name: "John Smith", Total: 123},
		{Name: "john smith", Total: 234},
		{Name: "John smoth", Total: 345},
	},
	FemaleTotals: []structs.Entry{
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
		nameList []string
	}
	tests := []struct {
		name              string
		args              args
		wantNamePositions []int
	}{
		{name: "Single Match", args: args{
			nameStr:  "John Smith",
			nameList: []string{"andrew smith", "dave smith", "John Smith", "maybe john smith"}},
			wantNamePositions: []int{2, 3},
		},
		{name: "Multiple Match (case insensitive)", args: args{
			nameStr:  "John Smith",
			nameList: []string{"john smith", "john not smith", "John Smith", "john Smith"}},
			wantNamePositions: []int{0, 2, 3},
		},
		{name: "No Match on Spelling", args: args{
			nameStr:  "John Smith",
			nameList: []string{"jim smof", "dof smith", "john smof"}},
			wantNamePositions: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNamePositions := NameSearch(tt.args.nameStr, &tt.args.nameList); !reflect.DeepEqual(gotNamePositions, tt.wantNamePositions) {
				t.Errorf("NameSearch() = %v, want %v", gotNamePositions, tt.wantNamePositions)
			}
		})
	}
}
