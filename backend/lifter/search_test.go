package lifter

import (
	"backend/structs"
	"reflect"
	"testing"
)

// todo: add more details to allow more strict testing
var sampleLeaderboardData = &structs.LeaderboardData{
	AllTotals: []*structs.Lift{
		{Lifter: &structs.Lifter{Name: "John Smith"}, Event: &structs.Event{}, Total: structs.NewFixedFloat(123)},
		{Lifter: &structs.Lifter{Name: "john smith"}, Event: &structs.Event{}, Total: structs.NewFixedFloat(234)},
		{Lifter: &structs.Lifter{Name: "John smoth"}, Event: &structs.Event{}, Total: structs.NewFixedFloat(345)},
		{Lifter: &structs.Lifter{Name: "Joanne Smith"}, Event: &structs.Event{}, Total: structs.NewFixedFloat(123)},
		{Lifter: &structs.Lifter{Name: "joanne smith"}, Event: &structs.Event{}, Total: structs.NewFixedFloat(234)},
		{Lifter: &structs.Lifter{Name: "joanne smith"}, Event: &structs.Event{}, Total: structs.NewFixedFloat(235)},
		{Lifter: &structs.Lifter{Name: "joanne Smoth"}, Event: &structs.Event{}, Total: structs.NewFixedFloat(345)},
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
		nameStr string
		roster  structs.LifterRoster
	}
	tests := []struct {
		name string
		args args
		want structs.NameSearchResults
	}{
		{name: "Single Match", args: args{
			nameStr: "Dave Smith",
			roster: structs.LifterRoster{Lifters: []*structs.Lifter{
				{Name: "andrew smith"},
				{Name: "Dave Smith"},
				{Name: "John Smith"},
				{Name: "maybe john smith"},
			}}},
			want: structs.NameSearchResults{
				Names: []structs.NameSearch{{NameStr: "Dave Smith"}},
				Total: 1,
			},
		},
		{name: "Multiple Match (case insensitive)", args: args{
			nameStr: "John Smith",
			roster: structs.LifterRoster{Lifters: []*structs.Lifter{
				{Name: "john smith"},
				{Name: "john not smith"},
				{Name: "John Smith"},
				{Name: "john Smith"},
			}}},
			want: structs.NameSearchResults{
				Names: []structs.NameSearch{{NameStr: "john smith"}, {NameStr: "John Smith"}, {NameStr: "john Smith"}},
				Total: 3,
			},
		},
		{name: "No Match on Spelling", args: args{
			nameStr: "John Smith",
			roster: structs.LifterRoster{Lifters: []*structs.Lifter{
				{Name: "jim smof"},
				{Name: "dof smith"},
				{Name: "john smof"},
			}}},
			want: structs.NameSearchResults{
				Names: []structs.NameSearch{{NameStr: ""}},
				Total: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NameSearch(tt.args.nameStr, tt.args.roster); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NameSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarNames(t *testing.T) {
	type args struct {
		nameSearch structs.NameSearch
		roster     structs.LifterRoster
	}
	tests := []struct {
		name                string
		args                args
		wantSimilaritySlice structs.NameSimilarityResults
	}{
		{name: "Single Match", args: args{
			nameSearch: structs.NameSearch{NameStr: "Chris Murray", Federation: "UK"},
			roster: structs.LifterRoster{Lifters: []*structs.Lifter{
				{Name: "Frankie Murray", PrimaryFederation: "US"},
				{Name: "Chris Murray", PrimaryFederation: "UK"},
				{Name: "MURRAY Chris", PrimaryFederation: "IWF"},
				{Name: "MURRAY Christopher John", PrimaryFederation: "IWF"},
			}}},
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
			got := SimilarNames(tt.args.nameSearch, &tt.args.roster)

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
