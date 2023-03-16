package events

import (
	"backend/structs"
	"reflect"
	"testing"
)

func TestFetchEvent(t *testing.T) {
	sampleLeaderboard := structs.LeaderboardData{
		MaleTotals: []structs.Entry{
			{Event: "A"},
			{Event: "B"},
			{Event: "C"},
			{Event: "D"},
			{Event: "E"},
			{Event: "F"},
		},
		FemaleTotals: []structs.Entry{
			{Event: "A"},
			{Event: "B"},
			{Event: "C"},
			{Event: "D"},
			{Event: "E"},
		},
	}
	type args struct {
		eventName   string
		leaderboard *structs.LeaderboardData
	}
	tests := []struct {
		name          string
		args          args
		wantEventData []structs.Entry
	}{
		{name: "FetchEventA", args: args{eventName: "A", leaderboard: &sampleLeaderboard}, wantEventData: []structs.Entry{{Event: "A"}, {Event: "A"}}},
		{name: "FetchEventF", args: args{eventName: "F", leaderboard: &sampleLeaderboard}, wantEventData: []structs.Entry{{Event: "F"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEventData := FetchEvent(tt.args.eventName, tt.args.leaderboard); !reflect.DeepEqual(gotEventData, tt.wantEventData) {
				t.Errorf("FetchEvent() = %v, want %v", gotEventData, tt.wantEventData)
			}
		})
	}
}
