package responders

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func init() {
	fmt.Println(addNonprefixListener("logmessage", logMessageIntoDB))
}

func logMessageIntoDB(message *discordgo.Message) (err error) {
	return jackal.Database.AddMessage(message)
}
