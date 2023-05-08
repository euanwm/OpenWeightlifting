package events

import "backend/structs"

// FetchEvent should use the exact string
func FetchEvent(eventName string, leaderboard *structs.LeaderboardData) (eventData []structs.Entry) {
	for _, lift := range leaderboard.AllTotals {
		if lift.Event == eventName {
			eventData = append(eventData, lift)
		}
	}
	return
}
