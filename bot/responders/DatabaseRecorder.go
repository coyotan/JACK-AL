package responders

import "github.com/bwmarrin/discordgo"

func init() {
	addNonprefixListener("", logMessageIntoDB)
}

func logMessageIntoDB(message *discordgo.Message) (err error) {
	return jackal.Database.AddMessage(message)
}
