package responders

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/txgruppi/parseargs-go"
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
	addCreateListener("ping", responderPong)
	addCreateListener("loaddb", responderLoadDB)
	addCreateListener("pulldb", responderPullDB)
}

func responderPong(message *discordgo.Message) (err error) {
	jackal.Logger.Console.Println("Discord message received: ", message.Content)

	if strings.ToLower(message.Content[len(jackal.Discord.CommandPrefix):]) == "ping" {
		_, err = jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "Pong!")
		jackal.Logger.Info.Println("Received ping, Pong!")
	}

	return
}

func responderLoadDB(message *discordgo.Message) (err error) {
	jackal.Logger.Console.Println("Received message with content: " + message.Content)

	args, err := parseargs.Parse(message.Content[len("!loaddb"):])

	if len(args) > 1 {
		err = jackal.DB.Put("root", args[0], args[1])

		if err != nil {
			jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "Something weird happened when we tried to put data into the database. We're crashing now!")
			fmt.Println(err)
		}
	}

	return
}

func responderPullDB(message *discordgo.Message) (err error) {
	jackal.Logger.Console.Println("Received message with content: " + message.Content)

	args, err := parseargs.Parse(message.Content[len("!pulldb"):])

	if len(args) >= 1 {
		val, err := jackal.DB.Get("root", args[0])

		if err != nil {
			jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "Something weird happened when we tried to put data into the database. We're crashing now!")
			fmt.Println(err)
		} else {
			jackal.Discord.Session.ChannelMessageSend(message.ChannelID, string(val))
		}
	}

	return
}
