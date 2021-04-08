package responders

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

//DOCUMENTATION:
/*
NO FILES INTENDING TO USE THE DISCORD INTERFACE MAY BE ALPHABETICALLY SUPERIOR TO 0_loadMods.go. This will cause load order issues.
Public Variables accessible from this location:
jackal - Core configuration.
jackal.Discord - Contains Discord configuration
*/

func init() {
	addCreateListener("reactionact", responderReactionAction)

	addReactListener("poll", responderPoll)

}

func responderReactionAction(message *discordgo.Message) (err error) {
	_, err = jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "This function has not yet been implemented.")
	fmt.Println(message.Content)
	return
}

func responderPoll(react *discordgo.MessageReaction) (err error) {
	_, err = jackal.Discord.Session.ChannelMessageSend(react.ChannelID, "This function has not yet been implemented.")

	return
}
