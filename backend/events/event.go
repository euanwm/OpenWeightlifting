package events

import "backend/structs"

// FetchEvent should use the exact string
func FetchEvent(eventName string, leaderboard *structs.LeaderboardData) (eventData []structs.Entry) {
	// todo: make this nicer, it's a temporary fix to get rid of some dumber shit I did
	for _, lift := range leaderboard.MaleTotals {
		if lift.Event == eventName {
			eventData = append(eventData, lift)
		}
	}
	for _, lift := range leaderboard.FemaleTotals {
		if lift.Event == eventName {
			eventData = append(eventData, lift)
		}
	}
	return
}
