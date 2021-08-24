package responders

import (
	"errors"
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

//reAction is a struct that can easily be marshaled into a json and back so that it can be stored in the databuckets. This allows for us to store reaction actions in a map that can be accessed by the messageID
type reAction struct {
	Reaction string `json:"reaction"`
	Action   string `json:"action"`
	Argument string `json:"argument"`
}

func init() {
	//	addCreateListener("reactionact", responderReactionAction)
	//TODO: Figure out how Vix was able to activate these functions despite them not being fully implemented.
	//	addReactListener("addreaction", responderPoll)
}

func responderReactionAction(message *discordgo.Message) (err error) {

	var watchMessageID string
	var actionWord = true
	var reaction = false

	var action string
	var argument string

	if len(message.Content) > 12 && strings.ToLower(message.Content[len(jackal.Discord.CommandPrefix):12]) == "reactionact" {
		//Parseargs in order: reactionact messageID Action1 Reaction Action2 Reaction Action3 Reaction [...]

		if len(message.Content) > 31 {
			//Get the messageID
			watchMessageID = message.Content[13:31]

			//TODO: Add reaction to the message at the ID listed here.

			//Process every third item from the
			args, err := parseargs.Parse(message.Content[31:])

			if err != nil {
				return err
			} else if len(args) < 3 {
				//Tell the user they fucked up here in a Discord message.
				return errors.New("not enough information provided to create a reaction action")
			}

			for index, value := range args {

				if index == 0 {
					actionWord = true
					reaction = false
				}

				if actionWord && !reaction {
					fmt.Println("Action is " + value)
					action = value
					actionWord = !actionWord
				} else if !actionWord && !reaction {
					fmt.Println("Argument is " + value)
					argument = value
					reaction = !reaction
				} else if !actionWord && reaction {
					fmt.Println("Reaction is " + value)
					//TODO: Add the message to the database with the action, argument, and reaction keywords.
					doWork(action, argument, value)
					return err
				}
			}

			fmt.Println(watchMessageID)

		} else { //If the user did not post something that passes the minimal length of a action and a reaction.

			return errors.New("not enough information provided to create a reaction action")
		}
		_, err = jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "We need to add more stuff for this to work!")
		fmt.Println(message.Content)
	}

	return
}

func doWork(action string, argument string, reaction string) (err error) {
	switch strings.ToLower(action) {
	//Add calendar scheduling support in the future, maybe?
	case "addrole":
		//Actually add their role.
	case "vote":
		//To be implemented at a very later date.
	case "setprefix":
		//Literally set the prefix for the username.
	case "setsuffix":
		//Literally set the suffix for the username
	case "void":
		//Do nothing.
	}
	return
}

func responderPoll(react *discordgo.MessageReaction) (err error) {
	_, err = jackal.Discord.Session.ChannelMessageSend(react.ChannelID, "This function has not yet been implemented.")

	//Get the messageID and actions from the databucket.
	//Pass the actions into doWork(action, argument, reaction)

	return
}
