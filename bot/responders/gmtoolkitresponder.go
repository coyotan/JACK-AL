package responders

import (
	"github.com/bwmarrin/discordgo"
	"github.com/txgruppi/parseargs-go"
	"math"
	"strconv"
)

//DOCUMENTATION:
/*
NO FILES INTENDING TO USE THE DISCORD INTERFACE MAY BE ALPHABETICALLY SUPERIOR TO 0_loadMods.go. This will cause load order issues.
Public Variables accessible from this location:
jackal - Core configuration.
jackal.Discord - Contains Discord configuration
*/
//The command we will use for the hypotenuse command.
var hypo = "hypo"

func init() {
	addCreateListener(hypo, hypotenuseResponder)
}

func hypotenuseResponder(message *discordgo.Message) (err error) {
	var a2, b2 int

	jackal.Logger.Console.Println("Received message with content: " + message.Content)

	args, err := parseargs.Parse(message.Content[len("!"+hypo):])

	//Because of the AllCast command fallthrough, sometimes this command responds when it really shouldn't.
	if len(args) > 1 && message.Content[len("!"+hypo):] == "!"+hypo {
		if a2, err = strconv.Atoi(args[0]); err != nil {
			jackal.Logger.Error.Println(err)
			jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "A had an error converting to an int!")
			return
		}

		if b2, err = strconv.Atoi(args[0]); err != nil {
			jackal.Logger.Error.Println(err)
			jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "B had an error converting to an int!")
			return
		}

		c := math.Sqrt(float64((a2 * a2) + (b2 + b2)))

		_, err = jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "The distance measured is "+strconv.FormatFloat(c, 'c', 2, 64))

	} else {
		//Respond that we don't have enough arguments. WE SHOULD NOT DO THIS IF WE WERE PART OF THE ALL CALL
		//jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "You did not provide enough arguments to execute the function. A and B must be provided to get C.")
	}

	return
}
