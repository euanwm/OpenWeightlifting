package structs

import (
	"backend/enum"
	"reflect"
	"testing"
)

func TestLift_WithinWeightClass(t *testing.T) {
	sampleLift := Lift{
		Bodyweight: NewFixedFloat(100),
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
				Upper:  NewFixedFloat(101),
				Lower:  NewFixedFloat(99),
			}},
			want: true,
		},
		{name: "CatchAll", args: args{
			gender: enum.Male,
			catData: WeightClass{
				Gender: enum.ALLCATS,
				Upper:  NewFixedFloat(101),
				Lower:  NewFixedFloat(99),
			}},
			want: true,
		},
		{name: "OutsideClass", args: args{
			gender: enum.Male,
			catData: WeightClass{
				Gender: enum.Male,
				Upper:  NewFixedFloat(99),
				Lower:  NewFixedFloat(98),
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleLift.WithinWeightClass(tt.args.gender, tt.args.catData); got != tt.want {
				t.Errorf("WithinWeightClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLift_WithinYear(t *testing.T) {
	sampleLift := Lift{
		Event: &Event{Date: "2020-01-01"},
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
			if got := sampleLift.WithinYear(tt.args.year); got != tt.want {
				t.Errorf("WithinYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLift_WithinDates(t *testing.T) {
	sampleLift := Lift{Event: &Event{Date: "2021-02-16"}}
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
			if got := sampleLift.WithinDates(tt.args.startDate, tt.args.endDate); got != tt.want {
				t.Errorf("WithinYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_SelectedFederation(t *testing.T) {
	sampleEvent := Event{
		Federation: "UK",
	}
	type args struct {
		fed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "SelectedFed", args: args{fed: "UK"}, want: true},
		{name: "NotSelectedFed", args: args{fed: "DrugsDrugsDrugs"}, want: false},
		{name: "AllFeds", args: args{fed: enum.ALLFEDS}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleEvent.SelectedFederation(tt.args.fed); got != tt.want {
				t.Errorf("SelectedFederation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeaderboardData_Select(t *testing.T) {
	sampleLeaderboard := LeaderboardData{
		AllTotals:    []*Lift{},
		AllSinclairs: []*Lift{},
	}
	type args struct {
		sortBy string
	}
	tests := []struct {
		name string
		args args
		want []*Lift
	}{
		{name: "SelectTotal", args: args{sortBy: enum.Total}, want: []*Lift{}},
		{name: "SelectSinclair", args: args{sortBy: enum.Sinclair}, want: []*Lift{}},
		{name: "NeitherMale", args: args{sortBy: "neither"}, want: []*Lift{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleLeaderboard.Select(tt.args.sortBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
