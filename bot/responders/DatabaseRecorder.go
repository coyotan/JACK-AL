package responders

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func init() {
	fmt.Println(addNonprefixListener("logmessage", logMessageIntoDB))
}

func logMessageIntoDB(message *discordgo.Message) (err error) {
	err = jackal.Database.AddUserFromMessage(message)
	if err != nil {
		return err
	}

	return jackal.Database.AddMessage(message)
}
