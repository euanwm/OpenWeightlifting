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
		Bodyweight: NewWeightKg(100),
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
				Upper:  NewWeightKg(101),
				Lower:  NewWeightKg(99),
			}},
			want: true,
		},
		{name: "CatchAll", args: args{
			gender: enum.Male,
			catData: WeightClass{
				Gender: enum.ALLCATS,
				Upper:  NewWeightKg(101),
				Lower:  NewWeightKg(99),
			}},
			want: true,
		},
		{name: "OutsideClass", args: args{
			gender: enum.Male,
			catData: WeightClass{
				Gender: enum.Male,
				Upper:  NewWeightKg(99),
				Lower:  NewWeightKg(98),
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
				Total:      NewWeightKg(100),
				BestSn:     NewWeightKg(40),
				BestCJ:     NewWeightKg(60),
				Bodyweight: NewWeightKg(50),
			},
			{
				Date:       "2020-01-02",
				Total:      NewWeightKg(200),
				BestSn:     NewWeightKg(80),
				BestCJ:     NewWeightKg(120),
				Bodyweight: NewWeightKg(100),
			},
			{
				Date:       "2020-01-03",
				Total:      NewWeightKg(300),
				BestSn:     NewWeightKg(120),
				BestCJ:     NewWeightKg(180),
				Bodyweight: NewWeightKg(150),
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

func TestLifterHistory_MakeRates(t *testing.T) {
	sampleLifterHistory := LifterHistory{
		Lifts: []Entry{
			{Sn1: NewWeightKg(55), Sn2: NewWeightKg(60), Sn3: NewWeightKg(70), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(-85), BestSn: NewWeightKg(70), BestCJ: NewWeightKg(80)},
			{Sn1: NewWeightKg(-55), Sn2: NewWeightKg(55), Sn3: NewWeightKg(-60), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(85), BestSn: NewWeightKg(55), BestCJ: NewWeightKg(85)},
			{Sn1: NewWeightKg(-60), Sn2: NewWeightKg(61), Sn3: NewWeightKg(-65), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(-85), BestSn: NewWeightKg(61), BestCJ: NewWeightKg(80)},
			{Sn1: NewWeightKg(58), Sn2: NewWeightKg(61), Sn3: NewWeightKg(-63), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(90), BestSn: NewWeightKg(61), BestCJ: NewWeightKg(90)},
			{Sn1: NewWeightKg(0), Sn2: NewWeightKg(0), Sn3: NewWeightKg(0), CJ1: NewWeightKg(0), CJ2: NewWeightKg(0), CJ3: NewWeightKg(0), BestSn: NewWeightKg(0), BestCJ: NewWeightKg(0)},
		},
	}
	tests := []struct {
		name string
		want []int
	}{
		{name: enum.Snatch, want: []int{50, 100, 25}},
		{name: enum.CleanAndJerk, want: []int{100, 0, 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleLifterHistory.MakeRates(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLifterHistory_BestLift(t *testing.T) {
	sampleLifterHistory := LifterHistory{
		Lifts: []Entry{
			{Sn1: NewWeightKg(55), Sn2: NewWeightKg(60), Sn3: NewWeightKg(70), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(-85), BestSn: NewWeightKg(70), BestCJ: NewWeightKg(80), Total: NewWeightKg(150)},
			{Sn1: NewWeightKg(-55), Sn2: NewWeightKg(55), Sn3: NewWeightKg(-60), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(85), BestSn: NewWeightKg(55), BestCJ: NewWeightKg(85), Total: NewWeightKg(140)},
			{Sn1: NewWeightKg(-60), Sn2: NewWeightKg(61), Sn3: NewWeightKg(-65), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(-85), BestSn: NewWeightKg(61), BestCJ: NewWeightKg(80), Total: NewWeightKg(141)},
			{Sn1: NewWeightKg(58), Sn2: NewWeightKg(61), Sn3: NewWeightKg(-63), CJ1: NewWeightKg(80), CJ2: NewWeightKg(-85), CJ3: NewWeightKg(90), BestSn: NewWeightKg(61), BestCJ: NewWeightKg(90), Total: NewWeightKg(151)},
		},
	}
	tests := []struct {
		name string
		want WeightKg
	}{
		{name: enum.Snatch, want: NewWeightKg(70)},
		{name: enum.CleanAndJerk, want: NewWeightKg(90)},
		{name: enum.Total, want: NewWeightKg(151)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleLifterHistory.BestLift(tt.name); got != tt.want {
				t.Errorf("BestLift() = %v, want %v", got, tt.want)
			}
		})
	}
}
