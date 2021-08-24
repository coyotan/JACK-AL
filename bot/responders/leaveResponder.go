package responders

import (
	"github.com/bwmarrin/discordgo"
	"github.com/coyotan/JACK-AL/botutils"
	"os"
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
	addCreateListener("leave", responderLeave)
}

func responderLeave(message *discordgo.Message) (err error) {
	if strings.ToLower(message.Content[len(jackal.Discord.CommandPrefix):]) == "leave" {
		jackal.Logger.Console.Println("Discord message received: ", message.Content)
		if isAdm, err := botutils.CheckAdminPermissions(jackal.Discord.Session, message); isAdm {
			if strings.ToLower(message.Content[len(jackal.Discord.CommandPrefix):]) == "leave" {
				jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "Leaving now!")
				jackal.Discord.Session.Close()
				jackal.Logger.Info.Println("Peacefully closing JACK-AL. Exit Code 100")
				os.Exit(100)
			}
		} else if err != nil {
			jackal.Logger.Error.Println("There was an error validating the permissions of user ", message.Author.Username+"#"+message.Author.Discriminator, " when they attempted to run the leave command.")
		} else {
			jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "You do not have permission to use that command")
			jackal.Logger.Warn.Println(message.Author.Username+"#"+message.Author.Discriminator, " attempted to run privileged command 'leave' but was rejected due to insufficient permissions.")
		}
	}
	return
}
