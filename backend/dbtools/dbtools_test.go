package dbtools

import (
	"backend/enum"
	"backend/structs"
	"reflect"
	"testing"
)

func TestBuildDatabase(t *testing.T) {
	t.Run("BuildDatabase", func(t *testing.T) {
		dbBuild := structs.LeaderboardData{}
		BuildDatabase(&dbBuild)
		if len(dbBuild.MaleTotals) == 0 {
			t.Errorf("BuildDatabase() = %v, want greater than 0", len(dbBuild.MaleTotals))
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
			if gotAllData := CollateAll(); !reflect.DeepEqual(reflect.TypeOf(gotAllData), reflect.TypeOf(tt.wantAllData)) {
				t.Errorf("CollateAll() = %v, want %v", reflect.TypeOf(gotAllData), reflect.TypeOf(tt.wantAllData))
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		bigData        []structs.Entry
		filterQuery    structs.LeaderboardPayload
		weightCat      structs.WeightClass
		lifterProfiles map[string]string
	}
	tests := []struct {
		name             string
		args             args
		wantFilteredData []structs.Entry
	}{
		// todo: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilteredData := Filter(tt.args.bigData, tt.args.filterQuery, tt.args.weightCat, tt.args.lifterProfiles); !reflect.DeepEqual(gotFilteredData, tt.wantFilteredData) {
				t.Errorf("Filter() = %v, want %v", gotFilteredData, tt.wantFilteredData)
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
			SortTotal(tt.args.sliceStructs)
		})
	}
}

func TestSortGender(t *testing.T) {
	type args struct {
		bigData [][]string
	}
	tests := []struct {
		name        string
		args        args
		wantMale    structs.AllData
		wantFemale  structs.AllData
		wantUnknown structs.AllData
	}{
		// todo: add test cases
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMale, gotFemale, gotUnknown := SortGender(tt.args.bigData)
			if !reflect.DeepEqual(gotMale, tt.wantMale) {
				t.Errorf("SortGender() gotMale = %v, want %v", gotMale, tt.wantMale)
			}
			if !reflect.DeepEqual(gotFemale, tt.wantFemale) {
				t.Errorf("SortGender() gotFemale = %v, want %v", gotFemale, tt.wantFemale)
			}
			if !reflect.DeepEqual(gotUnknown, tt.wantUnknown) {
				t.Errorf("SortGender() gotUnknown = %v, want %v", gotUnknown, tt.wantUnknown)
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
			if gotLineStruct := assignStruct(tt.args.line); !reflect.DeepEqual(gotLineStruct, tt.wantLineStruct) {
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
		{name: "FedDirs", wantFedDirs: []string{"AUS", "IWF", "UK", "US"}},
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
		{name: "InsertFederation", args: args{
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
			if gotAllEvents := loadAllFedEvents(tt.args.federation); !reflect.DeepEqual(reflect.TypeOf(gotAllEvents), reflect.TypeOf(tt.wantAllEvents)) {
				t.Errorf("loadAllFedEvents() = %v, want %v", gotAllEvents, reflect.TypeOf(tt.wantAllEvents))
			}
		})
	}
}

func Test_removeFollowingLifts(t *testing.T) {
	type args struct {
		bigData []structs.Entry
	}
	tests := []struct {
		name             string
		args             args
		wantFilteredData []structs.Entry
	}{
		{name: "RemoveFollowingLifts", args: args{bigData: []structs.Entry{{Name: "John Smith", Total: 100}, {Name: "John Smith", Total: 200}, {Name: "John Smith", Total: 300}}}, wantFilteredData: []structs.Entry{{Name: "John Smith", Total: 100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilteredData := removeFollowingLifts(tt.args.bigData); !reflect.DeepEqual(gotFilteredData, tt.wantFilteredData) {
				t.Errorf("removeFollowingLifts() = %v, want %v", gotFilteredData, tt.wantFilteredData)
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
			if gotGender := setGender(tt.args.entry); gotGender != tt.wantGender {
				t.Errorf("setGender() = %v, want %v", gotGender, tt.wantGender)
			}
		})
	}
}
