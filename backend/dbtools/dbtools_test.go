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
		var eventmetadata structs.EventsMetaData
		BuildDatabase(&dbBuild, &eventmetadata)
	}
}

func TestBuildDatabase(t *testing.T) {
	t.Run("BuildDatabase", func(t *testing.T) {
		dbBuild := structs.LeaderboardData{}
		var EventsData structs.EventsMetaData
		BuildDatabase(&dbBuild, &EventsData)
		if len(dbBuild.AllTotals) == 0 {
			t.Errorf("BuildDatabase() = %v, want greater than 0", len(dbBuild.AllTotals))
		}
	})
}

func TestCollateAll(t *testing.T) {
	tests := []struct {
		name        string
		wantAllData [][]string
	}{
		{name: "CollateAll", wantAllData: [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var eventmetadata structs.EventsMetaData
			if gotAllData := CollateAll(&eventmetadata); !reflect.DeepEqual(reflect.TypeOf(gotAllData), reflect.TypeOf(tt.wantAllData)) {
				t.Errorf("CollateAll() = %v, want %v", reflect.TypeOf(gotAllData), reflect.TypeOf(tt.wantAllData))
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
				bigData: []structs.Entry{{Date: "2023-06-01", Name: "John Smith", Total: 100, Federation: "BWL", Gender: enum.Male, Bodyweight: 109.00}, {Date: "2023-06-01", Name: "Dave Smith", Total: 200, Federation: "BWL", Gender: enum.Male, Bodyweight: 109.00}, {Date: "2023-06-01", Name: "Ethan Smith", Total: 300, Federation: "BWL", Gender: enum.Male, Bodyweight: 109.00}},
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
				Data: []structs.Entry{{Date: "2023-06-01", Name: "John Smith", Total: 100, Federation: "BWL", Gender: enum.Male, Bodyweight: 109.00}, {Date: "2023-06-01", Name: "Dave Smith", Total: 200, Federation: "BWL", Gender: enum.Male, Bodyweight: 109.00}, {Date: "2023-06-01", Name: "Ethan Smith", Total: 300, Federation: "BWL", Gender: enum.Male, Bodyweight: 109.00}},
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

func TestParseData(t *testing.T) {
	type args struct {
		bigData [][]string
	}
	tests := []struct {
		name        string
		args        args
		wantLifts   structs.AllData
		wantUnknown structs.AllData
	}{
		// todo: add test cases
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLifts, gotUnknown := ParseData(tt.args.bigData)
			if !reflect.DeepEqual(gotLifts, tt.wantLifts) {
				t.Errorf("ParseData() gotMale = %v, want %v", gotLifts, tt.wantLifts)
			}
			if !reflect.DeepEqual(gotUnknown, tt.wantUnknown) {
				t.Errorf("ParseData() gotUnknown = %v, want %v", gotLifts, tt.wantUnknown)
			}
		})
	}
}

func TestSortLiftsBy(t *testing.T) {
	type args struct {
		bigData []structs.Entry
		sortBy  string
	}
	tests := []struct {
		name          string
		args          args
		wantFinalData []structs.Entry
	}{
		{name: "SortBySinclair", args: args{bigData: []structs.Entry{{Sinclair: 300}, {Sinclair: 100}, {Sinclair: 200}}, sortBy: enum.Sinclair}, wantFinalData: []structs.Entry{{Sinclair: 300}, {Sinclair: 200}, {Sinclair: 100}}},
		{name: "SortByTotal", args: args{bigData: []structs.Entry{{Total: 300}, {Total: 100}, {Total: 200}}, sortBy: enum.Total}, wantFinalData: []structs.Entry{{Total: 300}, {Total: 200}, {Total: 100}}},
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
		sliceStructs []structs.Entry
		wantedSlice  []structs.Entry
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "NormalSort", args: args{sliceStructs: []structs.Entry{{Sinclair: 300}, {Sinclair: 100}, {Sinclair: 200}}, wantedSlice: []structs.Entry{{Sinclair: 100}, {Sinclair: 200}, {Sinclair: 300}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortTotal(tt.args.sliceStructs)
		})
	}
}

func TestSortTotal(t *testing.T) {
	type args struct {
		sliceStructs []structs.Entry
		wantedSlice  []structs.Entry
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "NormalSort", args: args{sliceStructs: []structs.Entry{{Total: 300}, {Total: 100}, {Total: 200}}, wantedSlice: []structs.Entry{{Total: 100}, {Total: 200}, {Total: 300}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortTotal(tt.args.sliceStructs)
		})
	}
}

func Test_assignStruct(t *testing.T) {
	var lineRaw = []string{
		"British U20 & U23 Weightlifting Championships 2017", "2017-10-01", "Men's Under 23 94Kg", "Edmon avetisyan", "93.8", "-146", "150", "-156", "180", "-190", "-192", "150", "180", "330", "BWL",
	}
	type args struct {
		line []string
	}
	tests := []struct {
		name           string
		args           args
		wantLineStruct structs.Entry
	}{
		{name: "AssignExpected", args: args{line: lineRaw}, wantLineStruct: structs.Entry{
			Event:      "British U20 & U23 Weightlifting Championships 2017",
			Date:       "2017-10-01",
			Gender:     "Men's Under 23 94Kg",
			Name:       "Edmon avetisyan",
			Bodyweight: 93.8,
			Sn1:        -146,
			Sn2:        150,
			Sn3:        -156,
			CJ1:        180,
			CJ2:        -190,
			CJ3:        -192,
			BestSn:     150,
			BestCJ:     180,
			Total:      330,
			Sinclair:   0,
			Federation: "BWL",
			Instagram:  "",
		},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLineStruct, _ := assignStruct(tt.args.line); !reflect.DeepEqual(gotLineStruct, tt.wantLineStruct) {
				t.Errorf("assignStruct() = %v, want %v", gotLineStruct, tt.wantLineStruct)
			}
		})
	}
}

func Test_getFedDirs(t *testing.T) {
	tests := []struct {
		name        string
		wantFedDirs []string
	}{
		{name: "FedDirs", wantFedDirs: []string{"AUS", "FFH", "IWF", "NVF", "OPEN", "UK", "US"}},
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
		name          string
		args          args
		wantAllEvents [][]string
	}{
		{name: "LoadAllEvents", args: args{federation: "UK"}, wantAllEvents: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var eventmetadata structs.EventsMetaData
			if gotAllEvents := loadAllFedEvents(tt.args.federation, &eventmetadata); !reflect.DeepEqual(reflect.TypeOf(gotAllEvents), reflect.TypeOf(tt.wantAllEvents)) {
				t.Errorf("loadAllFedEvents() = %v, want %v", gotAllEvents, reflect.TypeOf(tt.wantAllEvents))
			}
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
			if gotGender := getGender(tt.args.entry); gotGender != tt.wantGender {
				t.Errorf("getGender() = %v, want %v", gotGender, tt.wantGender)
			}
		})
	}
}

func Test_EventsMetaData_FetchEvent(t *testing.T) {
	var eventmetadata structs.EventsMetaData
	var leaderboarddata structs.LeaderboardData
	BuildDatabase(&leaderboarddata, &eventmetadata)
	federation, eventID := eventmetadata.FetchEventFP(0)
	println(federation, eventID)
}

func TestLoadSingleEvent(t *testing.T) {
	type args struct {
		federation string
		eventID    string
	}
	tests := []struct {
		name             string
		args             args
		wantTotalResults int
	}{
		{name: "LoadSingleEvent", args: args{federation: "AUS", eventID: "1000.csv"}, wantTotalResults: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEvent := len(LoadSingleEvent(tt.args.federation, tt.args.eventID)); !reflect.DeepEqual(gotEvent, tt.wantTotalResults) {
				t.Errorf("LoadSingleEvent() = %v, want %v", gotEvent, tt.wantTotalResults)
			}
		})
	}
}

func Test_EventsMetaData_FetchEventWithinDate(t *testing.T) {
	var eventmetadata structs.EventsMetaData
	var leaderboarddata structs.LeaderboardData
	BuildDatabase(&leaderboarddata, &eventmetadata)
	// expecting 368 events between the dates 2023-01-01 and 2023-05-01
	events := eventmetadata.FetchEventWithinDate("2023-01-01", "2023-05-01")
	reflect.DeepEqual(len(events), 368)
}
