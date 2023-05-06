package structs

import (
	"backend/enum"
	"reflect"
	"testing"
)

func TestAllData_ProcessNames(t *testing.T) {
	sampleAllData := AllData{
		Lifts: []Entry{{Name: "A"}, {Name: "B"}, {Name: "C"}, {Name: "D"}, {Name: "E"}},
	}
	tests := []struct {
		name      string
		wantNames []string
	}{
		{"Test ProcessNames", []string{"A", "B", "C", "D", "E"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNames := sampleAllData.ProcessNames(); !reflect.DeepEqual(gotNames, tt.wantNames) {
				t.Errorf("ProcessNames() = %v, want %v", gotNames, tt.wantNames)
			}
		})
	}
}

func TestEntry_WithinWeightClass(t *testing.T) {
	sampleEntry := Entry{
		Gender:     enum.Male,
		Bodyweight: 100,
	}
	type args struct {
		gender  string
		catData WeightClass
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "WithinClass", args: args{
			gender: enum.Male,
			catData: WeightClass{
				Gender: enum.Male,
				Upper:  101,
				Lower:  99,
			}},
			want: true,
		},
		{name: "CatchAll", args: args{
			gender: enum.Male,
			catData: WeightClass{
				Gender: enum.ALLCATS,
				Upper:  101,
				Lower:  99,
			}},
			want: true,
		},
		{name: "OutsideClass", args: args{
			gender: enum.Male,
			catData: WeightClass{
				Gender: enum.Male,
				Upper:  99,
				Lower:  98,
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleEntry.WithinWeightClass(tt.args.gender, tt.args.catData); got != tt.want {
				t.Errorf("WithinWeightClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntry_WithinYear(t *testing.T) {
	sampleEntry := Entry{
		Date: "2020-01-01",
	}
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "WithinYear", args: args{year: 2020}, want: true},
		{name: "OutsideYear", args: args{year: 2019}, want: false},
		{name: "AllYears", args: args{year: enum.AllYears}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleEntry.WithinYear(tt.args.year); got != tt.want {
				t.Errorf("WithinYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntry_SelectedFederation(t *testing.T) {
	sampleEntry := Entry{
		Federation: "BWL",
	}
	type args struct {
		fed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "SelectedFed", args: args{fed: "BWL"}, want: true},
		{name: "NotSelectedFed", args: args{fed: "DrugsDrugsDrugs"}, want: false},
		{name: "AllFeds", args: args{fed: enum.ALLFEDS}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleEntry.SelectedFederation(tt.args.fed); got != tt.want {
				t.Errorf("SelectedFederation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeaderboardData_FetchNames(t *testing.T) {
	sampleLeaderboard := LeaderboardData{
		AllNames: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
	}
	type args struct {
		posSlice []int
	}
	tests := []struct {
		name      string
		args      args
		wantNames []string
	}{
		{name: "FetchNamesMultiple", args: args{posSlice: []int{0, 1, 2, 3, 4}}, wantNames: []string{"A", "B", "C", "D", "E"}},
		{name: "FetchNamesSingle", args: args{posSlice: []int{0}}, wantNames: []string{"A"}},
		{name: "FetchNamesEmpty", args: args{posSlice: []int{}}, wantNames: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNames := sampleLeaderboard.FetchNames(tt.args.posSlice); !reflect.DeepEqual(gotNames, tt.wantNames) {
				t.Errorf("FetchNames() = %v, want %v", gotNames, tt.wantNames)
			}
		})
	}
}

func TestLeaderboardData_Query(t *testing.T) {
	sampleLeaderboard := LeaderboardData{
		AllNames:        []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
		MaleTotals:      []Entry{},
		FemaleTotals:    []Entry{},
		MaleSinclairs:   []Entry{},
		FemaleSinclairs: []Entry{},
	}
	type args struct {
		sortBy string
		gender string
	}
	tests := []struct {
		name string
		args args
		want []Entry
	}{
		{name: "TotalMale", args: args{sortBy: enum.Total, gender: enum.Male}, want: sampleLeaderboard.MaleTotals},
		{name: "TotalFemale", args: args{sortBy: enum.Total, gender: enum.Female}, want: sampleLeaderboard.FemaleTotals},
		{name: "SinclairMale", args: args{sortBy: enum.Sinclair, gender: enum.Male}, want: sampleLeaderboard.MaleSinclairs},
		{name: "SinclairFemale", args: args{sortBy: enum.Sinclair, gender: enum.Female}, want: sampleLeaderboard.FemaleSinclairs},
		{name: "NeitherMale", args: args{sortBy: "neither", gender: enum.Male}, want: []Entry{}},
		{name: "TotalNeither", args: args{sortBy: enum.Total, gender: "neither"}, want: []Entry{}},
		{name: "TotalNeither", args: args{sortBy: enum.Sinclair, gender: "neither"}, want: []Entry{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleLeaderboard.Query(tt.args.sortBy, tt.args.gender); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLifterHistory_GenerateChartData(t *testing.T) {
	sampleLifterHistory := LifterHistory{
		NameStr: "A",
		Lifts: []Entry{
			{
				Date:       "2020-01-01",
				Total:      100,
				BestSn:     40,
				BestCJ:     60,
				Bodyweight: 50,
			},
			{
				Date:       "2020-01-02",
				Total:      200,
				BestSn:     80,
				BestCJ:     120,
				Bodyweight: 100,
			},
			{
				Date:       "2020-01-03",
				Total:      300,
				BestSn:     120,
				BestCJ:     180,
				Bodyweight: 150,
			},
		},
	}
	tests := []struct {
		name string
		want ChartData
	}{
		{name: "GenerateChartData", want: ChartData{
			Dates: []string{"2020-01-01", "2020-01-02", "2020-01-03"},
			SubData: []ChartSubData{
				{
					Title:     "Competition Total",
					DataSlice: []float32{100, 200, 300},
				},
				{
					Title:     "Best Snatch",
					DataSlice: []float32{40, 80, 120},
				},
				{
					Title:     "Best C&J",
					DataSlice: []float32{60, 120, 180},
				},
				{
					Title:     "Bodyweight",
					DataSlice: []float32{50, 100, 150},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleLifterHistory.GenerateChartData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateChartData() = %v, want %v", got, tt.want)
			}
		})
	}
}
