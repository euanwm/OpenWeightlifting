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

func TestEntry_WithinDates(t *testing.T) {
	sampleEntry := Entry{Date: "2021-02-16"}
	type args struct {
		startDate string
		endDate   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "WithinDate", args: args{startDate: "2021-02-14", endDate: "2021-02-17"}, want: true},
		{name: "OutsideDate", args: args{startDate: "2019-01-01", endDate: "2019-12-31"}, want: false},
		{name: "AllDates", args: args{startDate: enum.ZeroDate, endDate: enum.MaxDate}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleEntry.WithinDates(tt.args.startDate, tt.args.endDate); got != tt.want {
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
		AllTotals: []Entry{
			{Name: "A"},
			{Name: "B"},
			{Name: "C"},
			{Name: "D"},
			{Name: "E"},
			{Name: "F"},
			{Name: "G"},
			{Name: "H"},
			{Name: "I"},
			{Name: "J"},
		},
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

func TestLeaderboardData_Select(t *testing.T) {
	sampleLeaderboard := LeaderboardData{
		AllTotals:    []Entry{},
		AllSinclairs: []Entry{},
	}
	type args struct {
		sortBy string
	}
	tests := []struct {
		name string
		args args
		want []Entry
	}{
		{name: "SelectTotal", args: args{sortBy: enum.Total}, want: sampleLeaderboard.AllTotals},
		{name: "SelectSinclair", args: args{sortBy: enum.Sinclair}, want: sampleLeaderboard.AllSinclairs},
		{name: "NeitherMale", args: args{sortBy: "neither"}, want: []Entry{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleLeaderboard.Select(tt.args.sortBy); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
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

func TestEntry_MakePercentage(t *testing.T) {
	sampleEntry := Entry{Sn1: 77.0, Sn2: -80.00, Sn3: 81, CJ1: 100, CJ2: 110, CJ3: 115}
	type args struct {
		lift string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "MakePercentageSnatch", args: args{lift: enum.Snatch}, want: 66},
		{name: "MakePercentageCJ", args: args{lift: enum.CleanAndJerk}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleEntry.MakePercentage(tt.args.lift); got != tt.want {
				t.Errorf("MakePercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}
