package responders

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

//DOCUMENTATION:
/*
NO FILES INTENDING TO USE THE DISCORD INTERFACE MAY BE ALPHABETICALLY SUPERIOR TO 0_loadMods.go. This will cause load order issues.
Public Variables accessible from this location:
jackal - Core configuration.
jackal.Discord - Contains Discord configuration
*/

func init() {
	addCreateListener("poll", responderPoll)
}

func responderPoll(message *discordgo.Message) (err error) {
	jackal.Logger.Console.Println("Discord message received: ", message.Content)

	if strings.ToLower(message.Content[len(jackal.Discord.CommandPrefix):]) == "ping" {
		_, err = jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "What kind of poll do you want to make?!")
	}

	return
}
