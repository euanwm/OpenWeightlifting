package lifter

import (
	"backend/structs"
	"reflect"
	"testing"
)

// todo: add more details to allow more strict testing
var sampleLeaderboardData = &structs.LeaderboardData{
	AllTotals: []*structs.Lift{
		{Lifter: &structs.Lifter{Name: "John Smith"}, Event: &structs.Event{}, Total: structs.NewWeightKg(123)},
		{Lifter: &structs.Lifter{Name: "john smith"}, Event: &structs.Event{}, Total: structs.NewWeightKg(234)},
		{Lifter: &structs.Lifter{Name: "John smoth"}, Event: &structs.Event{}, Total: structs.NewWeightKg(345)},
		{Lifter: &structs.Lifter{Name: "Joanne Smith"}, Event: &structs.Event{}, Total: structs.NewWeightKg(123)},
		{Lifter: &structs.Lifter{Name: "joanne smith"}, Event: &structs.Event{}, Total: structs.NewWeightKg(234)},
		{Lifter: &structs.Lifter{Name: "joanne smith"}, Event: &structs.Event{}, Total: structs.NewWeightKg(235)},
		{Lifter: &structs.Lifter{Name: "joanne Smoth"}, Event: &structs.Event{}, Total: structs.NewWeightKg(345)},
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
			wantNameSlice: []string{""},
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

func TestSimilarNames(t *testing.T) {
	type args struct {
		nameSearch structs.NameSearch
		nameList   []structs.Entry
	}
	tests := []struct {
		name                string
		args                args
		wantSimilaritySlice structs.NameSimilarityResults
	}{
		{name: "Single Match", args: args{
			nameSearch: structs.NameSearch{NameStr: "Chris Murray", Federation: "UK"},
			nameList: []structs.Entry{
				{Name: "Frankie Murray", Federation: "US"},
				{Name: "Chris Murray", Federation: "UK"},
				{Name: "MURRAY Chris", Federation: "IWF"},
				{Name: "MURRAY Christopher John", Federation: "IWF"},
			}},
			// Expected: all 4 entries score above the threshold (0.6).
			// Sorted by descending score:
			//   "Chris Murray" UK  → 1.0  (exact)
			//   "MURRAY Chris" IWF → 1.0  (token sort neutralises format flip)
			//   "MURRAY Christopher John" IWF → ~0.78 (token JW: "Chris"≈"Christopher")
			//   "Frankie Murray" US → ~0.70 (shared surname)
			wantSimilaritySlice: structs.NameSimilarityResults{
				Names: []structs.NameSimilarity{
					{NameStr: "Chris Murray", Federation: "UK", Score: 1},
					{NameStr: "MURRAY Chris", Federation: "IWF", Score: 1},
					{NameStr: "MURRAY Christopher John", Federation: "IWF", Score: 0.7759684},
					{NameStr: "Frankie Murray", Federation: "US", Score: 0.69714284},
				},
				Total: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimilarNames(tt.args.nameSearch, &tt.args.nameList)

			if got.Total != tt.wantSimilaritySlice.Total {
				t.Errorf("SimilarNames() Total = %d, want %d", got.Total, tt.wantSimilaritySlice.Total)
			}
			if len(got.Names) != len(tt.wantSimilaritySlice.Names) {
				t.Fatalf("SimilarNames() len(Names) = %d, want %d", len(got.Names), len(tt.wantSimilaritySlice.Names))
			}
			for i, want := range tt.wantSimilaritySlice.Names {
				g := got.Names[i]
				if g.NameStr != want.NameStr || g.Federation != want.Federation || g.Score != want.Score {
					t.Errorf("SimilarNames() Names[%d] = {%q, %q, %v}, want {%q, %q, %v}",
						i, g.NameStr, g.Federation, g.Score,
						want.NameStr, want.Federation, want.Score)
				}
			}
			// Verify results are in descending score order.
			for i := 1; i < len(got.Names); i++ {
				if got.Names[i].Score > got.Names[i-1].Score {
					t.Errorf("SimilarNames() scores not descending at index %d: %v > %v",
						i, got.Names[i].Score, got.Names[i-1].Score)
				}
			}
		})
	}
}
