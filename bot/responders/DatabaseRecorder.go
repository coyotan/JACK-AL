package responders

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func init() {
	addNonprefixListener("logmessage", logMessageIntoDB)
	addCreateListener("log", messageLogDispatcher)
}

func messageLogDispatcher(message *discordgo.Message) (err error) {

	if len(message.Content) > 1 {

		command := strings.Fields(strings.ToLower(message.Content))

		switch command[1] {

		case "get":

			switch command[2] {
			case "history":
				//!log get history 50
				if len(command) > 2 {
					fmt.Println("Log get history with argument")
					conv, err := strconv.Atoi(command[3])
					if err != nil {
						break
					}

					err = logHistoryIntoDB(message, conv)

				} else {
					//!log get history
					fmt.Println("Log get history without argument")
					err = logHistoryIntoDB(message, 100)
				}
			default:
			}

		default:
			fmt.Println("Invalid or nil argument. !log. Eg. '!log get history'")
		}
	}

	return err
}

func logMessageIntoDB(message *discordgo.Message) (err error) {
	err = jackal.Database.AddUserFromMessage(message)
	if err != nil {
		return err
	}

	return jackal.Database.AddMessage(message)
}

func logHistoryIntoDB(message *discordgo.Message, amount int) (err error) {
	var lastID = message.ID

	if amount > 100 {
		//If the user provided a number that was greater than 100.
		loopTime, remainder := divMod(amount, 100)

		for i := 0; i < loopTime; i++ {

			fmt.Println("Debug: ", i*100)

			var messages, err = jackal.Discord.Session.ChannelMessages(message.ChannelID, 100, lastID, "", "")
			if err != nil {
				return err
			}

			for j := 0; j < len(messages); j++ {
				logMessageIntoDB(messages[j])

				if err != nil {
					return err
				}
			}
			//Set the lastID to the new lastID
			fmt.Println(len(messages))
			lastID = messages[len(messages)-1].ID
		}

		if remainder > 1 {
			messages, err := jackal.Discord.Session.ChannelMessages(message.ChannelID, remainder, lastID, "", "")
			if err != nil {
				return err
			}

			for k := 0; k < len(messages); k++ {
				err = logMessageIntoDB(messages[k])
				if err != nil {
					return err
				}
			}
		}
	} else {
		//If the user provided a number that was less than 100.
		messages, err := jackal.Discord.Session.ChannelMessages(message.ChannelID, amount, lastID, "", "")
		if err != nil {
			return err
		}

		for k := 0; k < len(messages); k++ {
			err = logMessageIntoDB(messages[k])
			if err != nil {
				return err
			}
		}
	}

	return err
}

func divMod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}
