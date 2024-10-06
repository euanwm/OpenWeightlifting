package discordbot

import "github.com/bwmarrin/discordgo"

// DiscordBot is the main struct for the discord bot
type DiscordBot struct {
	Session         *discordgo.Session
	Channel         string
	PlatformChannel string
	// todo: add mutex synchronization
}

// New creates a new DiscordBot
func New(token string) (DiscordBot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return DiscordBot{}, err
	}
	return DiscordBot{Session: session}, nil
}

// OpenConnection Open opens a connection to Discord
func (d *DiscordBot) OpenConnection() error {
	return d.Session.Open()
}

// CloseConnection Close closes the connection to Discord
func (d *DiscordBot) CloseConnection() error {
	return d.Session.Close()
}

// PostMessage posts a message to a channel
func (d *DiscordBot) PostMessage(message string) (*discordgo.Message, error) {
	return d.Session.ChannelMessageSend(d.Channel, message)
}

// PostPlatformData because I haven't bothered to create a better way of handling multiple channels yet
func (d *DiscordBot) PostPlatformData(message string) (*discordgo.Message, error) {
	return d.Session.ChannelMessageSend(d.PlatformChannel, message)
}
