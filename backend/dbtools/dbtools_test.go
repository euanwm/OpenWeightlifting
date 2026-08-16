package dbtools

import (
	"backend/enum"
	"backend/structs"
	"reflect"
	"testing"
)

func BenchmarkBuildDatabase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dbBuild := structs.LeaderboardData{}
		var eventsData structs.EventsData
		var roster structs.LifterRoster
		BuildDatabase(&dbBuild, &eventsData, &roster)
	}
}

func TestBuildDatabase(t *testing.T) {
	t.Run("BuildDatabase", func(t *testing.T) {
		dbBuild := structs.LeaderboardData{}
		var eventsData structs.EventsData
		var roster structs.LifterRoster
		BuildDatabase(&dbBuild, &eventsData, &roster)
		if len(dbBuild.AllTotals) == 0 {
			t.Errorf("BuildDatabase() = %v, want greater than 0", len(dbBuild.AllTotals))
		}
	})
}

func TestCollateAll(t *testing.T) {
	tests := []struct {
		name         string
		wantAllLifts structs.AllLifts
	}{
		{name: "CollateAll", wantAllLifts: structs.AllLifts{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var eventsData structs.EventsData
			var roster structs.LifterRoster
			if gotAllLifts := CollateAll(&eventsData, &roster); !reflect.DeepEqual(reflect.TypeOf(gotAllLifts), reflect.TypeOf(tt.wantAllLifts)) {
				t.Errorf("CollateAll() = %v, want %v", reflect.TypeOf(gotAllLifts), reflect.TypeOf(tt.wantAllLifts))
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		bigData     []structs.Entry
		filterQuery structs.LeaderboardPayload
		weightCat   string
	}
	tests := []struct {
		name             string
		args             args
		wantFilteredData structs.LeaderboardResponse
	}{
		{
			name: "FilterByFederation",
			args: args{
				bigData: []structs.Entry{
					{Date: "2023-06-01", Name: "John Smith", Total: structs.NewWeightKg(100), Federation: "BWL", Gender: enum.Male, Bodyweight: structs.NewWeightKg(109.00)},
					{Date: "2023-06-01", Name: "Dave Smith", Total: structs.NewWeightKg(200), Federation: "BWL", Gender: enum.Male, Bodyweight: structs.NewWeightKg(109.00)},
					{Date: "2023-06-01", Name: "Ethan Smith", Total: structs.NewWeightKg(300), Federation: "BWL", Gender: enum.Male, Bodyweight: structs.NewWeightKg(109.00)},
				},
				filterQuery: structs.LeaderboardPayload{
					Start:       0,
					Stop:        10,
					SortBy:      enum.Total,
					Federation:  enum.ALLFEDS,
					WeightClass: "MALL",
					Year:        "69",
					StartDate:   "2023-01-01",
					EndDate:     "2024-01-01",
				},
				weightCat: "MALL",
			},
			wantFilteredData: structs.LeaderboardResponse{
				Size: 3,
				Data: []structs.Entry{
					{Date: "2023-06-01", Name: "John Smith", Total: structs.NewWeightKg(100), Federation: "BWL", Gender: enum.Male, Bodyweight: structs.NewWeightKg(109.00)},
					{Date: "2023-06-01", Name: "Dave Smith", Total: structs.NewWeightKg(200), Federation: "BWL", Gender: enum.Male, Bodyweight: structs.NewWeightKg(109.00)},
					{Date: "2023-06-01", Name: "Ethan Smith", Total: structs.NewWeightKg(300), Federation: "BWL", Gender: enum.Male, Bodyweight: structs.NewWeightKg(109.00)},
				},
			},
		},
	}
	var cache QueryCache
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilteredData := FilterLifts(tt.args.bigData, tt.args.filterQuery, WeightClassList[tt.args.weightCat], &cache); !reflect.DeepEqual(gotFilteredData, tt.wantFilteredData) {
				t.Errorf("FilterLifts() = %v, want %v", gotFilteredData, tt.wantFilteredData)
			}
		})
	}
}

func TestSortDate(t *testing.T) {
	type args struct {
		sliceStructs []structs.Entry
		wantedSlice  []structs.Entry
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "NormalSort", args: args{sliceStructs: []structs.Entry{{Date: "2020-04-16"}, {Date: "2021-03-18"}, {Date: "2019-08-24"}}, wantedSlice: []structs.Entry{{Date: "2019-08-24"}, {Date: "2020-04-16"}, {Date: "2021-03-18"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortDate(tt.args.sliceStructs)
			if !reflect.DeepEqual(tt.args.sliceStructs, tt.args.wantedSlice) {
				t.Errorf("SortDate() = %v, want %v", tt.args.sliceStructs, tt.args.wantedSlice)
			}
		})
	}
}

func TestSortLiftsBy(t *testing.T) {
	type args struct {
		bigData []*structs.Lift
		sortBy  string
	}
	tests := []struct {
		name          string
		args          args
		wantFinalData []*structs.Lift
	}{
		{name: "SortBySinclair", args: args{bigData: []*structs.Lift{{Sinclair: 300}, {Sinclair: 100}, {Sinclair: 200}}, sortBy: enum.Sinclair}, wantFinalData: []*structs.Lift{{Sinclair: 300}, {Sinclair: 200}, {Sinclair: 100}}},
		{name: "SortByTotal", args: args{bigData: []*structs.Lift{{Total: structs.NewWeightKg(300)}, {Total: structs.NewWeightKg(100)}, {Total: structs.NewWeightKg(200)}}, sortBy: enum.Total}, wantFinalData: []*structs.Lift{{Total: structs.NewWeightKg(300)}, {Total: structs.NewWeightKg(200)}, {Total: structs.NewWeightKg(100)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFinalData := SortLiftsBy(tt.args.bigData, tt.args.sortBy); !reflect.DeepEqual(gotFinalData, tt.wantFinalData) {
				t.Errorf("SortLiftsBy() = %v, want %v", gotFinalData, tt.wantFinalData)
			}
		})
	}
}

func TestSortSinclair(t *testing.T) {
	type args struct {
		sliceStructs []*structs.Lift
		wantedSlice  []*structs.Lift
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "NormalSort", args: args{sliceStructs: []*structs.Lift{{Sinclair: 300}, {Sinclair: 100}, {Sinclair: 200}}, wantedSlice: []*structs.Lift{{Sinclair: 100}, {Sinclair: 200}, {Sinclair: 300}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortTotal(tt.args.sliceStructs)
		})
	}
}

func TestSortTotal(t *testing.T) {
	type args struct {
		sliceStructs []*structs.Lift
		wantedSlice  []*structs.Lift
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "NormalSort", args: args{sliceStructs: []*structs.Lift{
			{Total: structs.NewWeightKg(300)},
			{Total: structs.NewWeightKg(100)},
			{Total: structs.NewWeightKg(200)},
		}, wantedSlice: []*structs.Lift{
			{Total: structs.NewWeightKg(100)},
			{Total: structs.NewWeightKg(200)},
			{Total: structs.NewWeightKg(300)},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortTotal(tt.args.sliceStructs)
		})
	}
}

func Test_getFedDirs(t *testing.T) {
	tests := []struct {
		name        string
		wantFedDirs []string
	}{
		{name: "FedDirs", wantFedDirs: []string{"AUS", "CH", "FFH", "IRE", "IWF", "NVF", "OPEN", "UK", "US"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFederationDirs := getFedDirs(); !reflect.DeepEqual(gotFederationDirs, tt.wantFedDirs) {
				t.Errorf("getFedDirs() = %v, want %v", gotFederationDirs, tt.wantFedDirs)
			}
		})
	}
}

func Test_insertFederation(t *testing.T) {
	type args struct {
		event      [][]string
		federation string
	}
	tests := []struct {
		name             string
		args             args
		wantNewEventData [][]string
	}{
		{name: "insertFederation", args: args{
			event: [][]string{{
				"British U20 & U23 Weightlifting Championships 2017", "2017-10-01", "Men's Under 23 94Kg", "Edmon avetisyan", "93.8", "-146", "150", "-156", "180", "-190", "-192", "150", "180", "330"}},
			federation: "UK",
		}, wantNewEventData: [][]string{{
			"British U20 & U23 Weightlifting Championships 2017", "2017-10-01", "Men's Under 23 94Kg", "Edmon avetisyan", "93.8", "-146", "150", "-156", "180", "-190", "-192", "150", "180", "330", "UK"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewEventData := insertFederation(tt.args.event, tt.args.federation); !reflect.DeepEqual(gotNewEventData, tt.wantNewEventData) {
				t.Errorf("insertFederation() = %v, want %v", gotNewEventData, tt.wantNewEventData)
			}
		})
	}
}

func Test_loadAllFedEvents(t *testing.T) {
	type args struct {
		federation string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "LoadUKEvents", args: args{federation: "UK"}},
		{name: "LoadNVFEvents", args: args{federation: "NVF"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var eventsData structs.EventsData
			var allLifts structs.AllLifts
			var roster structs.LifterRoster
			loadAllFedEvents(tt.args.federation, &eventsData, &allLifts, &roster)
		})
	}
}

func Test_setGender(t *testing.T) {
	type args struct {
		entry *structs.Entry
	}
	tests := []struct {
		name       string
		args       args
		wantGender string
	}{
		{name: "DirectMatchMale", args: args{entry: &structs.Entry{Gender: enum.Male}}, wantGender: enum.Male},
		{name: "DirectMatchFemale", args: args{entry: &structs.Entry{Gender: enum.Female}}, wantGender: enum.Female},
		{name: "ContainsMen", args: args{entry: &structs.Entry{Gender: "Men's"}}, wantGender: enum.Male},
		{name: "ContainsWomen", args: args{entry: &structs.Entry{Gender: "Women's"}}, wantGender: enum.Female},
		{name: "CatchUnknown", args: args{entry: &structs.Entry{Gender: "something else"}}, wantGender: enum.Unknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGender := GetGender(tt.args.entry); gotGender != tt.wantGender {
				t.Errorf("GetGender() = %v, want %v", gotGender, tt.wantGender)
			}
		})
	}
}

func Test_CreateSingleEvent(t *testing.T) {
	var events structs.EventsData
	var roster structs.LifterRoster
	var lifts structs.AllLifts
	t.Run("CreateSingleEvent", func(t *testing.T) {
		createSingleEvent("AUS", "1000.csv", &events, &lifts, &roster)
		if len(events.Events) != 1 {
			t.Errorf("CreateSingleEvent() = %v, want 1", len(events.Events))
		}
		if len(lifts.Lifts) != 18 {
			t.Errorf("CreateSingleEvent() = %v, want 18", len(lifts.Lifts))
		}
		// check that the lifts have referenced lifters
		for _, lift := range lifts.Lifts {
			if lift.Lifter == nil {
				t.Errorf("CreateSingleEvent() = %v, want referenced lifter", lift.Lifter)
			}
		}
	})
}
