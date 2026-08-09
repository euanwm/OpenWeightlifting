package events

import (
	"backend/structs"
	"reflect"
	"testing"
)

func TestFetchEvent(t *testing.T) {
	sampleLeaderboard := structs.LeaderboardData{
		AllTotals: []*structs.Lift{
			{Lifter: &structs.Lifter{}, Event: &structs.Event{Name: "A"}},
			{Lifter: &structs.Lifter{}, Event: &structs.Event{Name: "A"}},
			{Lifter: &structs.Lifter{}, Event: &structs.Event{Name: "B"}},
			{Lifter: &structs.Lifter{}, Event: &structs.Event{Name: "C"}},
			{Lifter: &structs.Lifter{}, Event: &structs.Event{Name: "D"}},
			{Lifter: &structs.Lifter{}, Event: &structs.Event{Name: "E"}},
			{Lifter: &structs.Lifter{}, Event: &structs.Event{Name: "F"}},
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
				t.Errorf("FetchEventFP() = %v, want %v", gotEventData, tt.wantEventData)
			}
		})
	}
}
