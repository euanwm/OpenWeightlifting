package feedbackbot

import "github.com/bwmarrin/discordgo"

const (
	// FeedbackChannelID - The channel ID for the feedback channel
	FeedbackChannelID = "1154508323241070632"
)

// FeedbackBot - Struct for the feedback bot
type FeedbackBot struct {
	Session *discordgo.Session
}

// New - Returns a new FeedbackBot
func New(token string) (*FeedbackBot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &FeedbackBot{Session: session}, nil
}

// PostToChannel - Posts a message to a channel
func (fb *FeedbackBot) PostToChannel(channelID string, message string) error {
	_, err := fb.Session.ChannelMessageSend(channelID, message)
	if err != nil {
		return err
	}
	return nil
}
