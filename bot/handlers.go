package bot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

//Function for receiving and processing the discord ready message once it is sent from the API.
func Ready(s *discordgo.Session, r *discordgo.Ready) {
	jackal.Logger.Info.Println("Discord Ready Message Received. Username:", r.User.Username, " User ID:", r.User.ID)
	jackal.Discord.User = r.User
	jackal.Discord.Session = s
	jackal.Logger.Info.Println("Discord Session data has been initialized.")
}

func createDispatch(_ *discordgo.Session, created *discordgo.MessageCreate) {
	if created.Author.ID != jackal.Discord.User.ID {
		var totalListeners= 0

		fields := strings.Fields(strings.ToLower(created.Message.Content))

		if string(fields[0][0]) == jackal.Discord.CommandPrefix {
			if val, ok := jackal.Discord.CreateListeners[fields[0][1:]]; ok {
				err := val(created.Message)

				if err != nil {
					jackal.Logger.Error.Println("Responder is 10-33", err)
				}

				totalListeners++
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

				jackal.Logger.Console.Println("Dispatched to ", totalListeners, " listeners. All responders are 10-8.")
			}
		}
	}
}