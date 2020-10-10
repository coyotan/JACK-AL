package bot

import (
	"github.com/CoyoTan/JACK-AL/structs"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func ready(s *discordgo.Session, r *discordgo.Ready) {
	jackal.Logger.Info.Println("Discord Ready Message Received. Username:", r.User.Username, " User ID:", r.User.ID)
	jackal.Discord.User = r.User
	jackal.Discord.Session = s
	jackal.Logger.Info.Println("Discord Session data has been initialized.")
}

func createDispatch(_ *discordgo.Session, created *discordgo.MessageCreate) {
	if created.Author.ID != jackal.Discord.User.ID {
		var totalListeners = 0

		//Add the message to the cache for logging, if it's deleted soon!
		jackal.Discord.Session.State.MessageAdd(created.Message)

		fields := strings.Fields(strings.ToLower(created.Message.Content))

		if string(fields[0][0]) == jackal.Discord.CommandPrefix {
			if val, ok := jackal.Discord.CreateListeners[fields[0][1:]]; ok {
				err := val(created.Message)

				if err != nil {
					jackal.Logger.Error.Println("Responder is 10-33", err)
				}

				totalListeners = 1
			} else {
				//If we cannot find the specific command we are looking for, tell EVERYONE what we found...
				for _, v := range jackal.Discord.CreateListeners {
					err := v(created.Message)

					if err != nil {
						jackal.Logger.Error.Println("Responder is 10-33", err)
					} else {
						totalListeners++
					}
				}
			}

			jackal.Logger.Console.Println("Dispatched to ", totalListeners, " listeners. All responders are 10-8.")
		}
	}
}

//Thanks to the new "BeforeDelete" method, this is all this function really needs to be! It's so much simpler now!
func deleteDispatch(_ *discordgo.Session, deleted *discordgo.MessageDelete) {
	if deleted.Author.ID != jackal.Discord.User.ID {
		var totalListeners = 0
		//If we cannot find the specific command we are looking for, tell EVERYONE what we found...
		for _, v := range jackal.Discord.DeleteListeners {
			err := v(deleted)

			if err != nil {
				jackal.Logger.Error.Println("Responder is 10-33", err)
			} else {
				totalListeners++
			}
		}

		jackal.Logger.Console.Println("Dispatched to ", totalListeners, " listeners. All responders are 10-8.")
	}
}

//Thanks to the new "BeforeUpdate" method, this is all this function really needs to be! It's so much simpler now!
func editDispatch(_ *discordgo.Session, updated *discordgo.MessageUpdate) {
	if updated.Author.ID != jackal.Discord.User.ID {
		var totalListeners = 0
		//If we cannot find the specific command we are looking for, tell EVERYONE what we found...
		for _, v := range jackal.Discord.EditListeners {
			err := v(updated)

			if err != nil {
				jackal.Logger.Error.Println("Responder is 10-33", err)
			} else {
				totalListeners++
			}
		}

		jackal.Logger.Console.Println("Dispatched to ", totalListeners, " listeners. All responders are 10-8.")
	}
}

func initModDispatch(jackal *structs.CoreCfg) {
	var totalListeners = 0
	//If we cannot find the specific command we are looking for, tell EVERYONE what we found...
	for _, v := range jackal.Discord.InitModListeners {
		err := v(jackal)

		if err != nil {
			jackal.Logger.Error.Println("Responder is 10-33", err)
		} else {
			totalListeners++
		}
	}

	jackal.Logger.Console.Println("Dispatched to ", totalListeners, " listeners. All responders are 10-8.")
}
