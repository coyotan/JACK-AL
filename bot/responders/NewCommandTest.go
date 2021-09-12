package responders

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/coyotan/JACK-AL/structs"
	"strings"
)

var (
	MyGroup = structs.CommandGroup{
		Name: "MyTestGroup",
	}
)

func init() {
	MyGroup.NewCommand("ping1", newPingResponder, structs.Enabled+structs.Discord+structs.Guild)
	fmt.Println(MyGroup.RegisterAllCommands())
}

func newPingResponder(_ *discordgo.Session, created *discordgo.MessageCreate) (err error) {

	message := created

	jackal.Logger.Console.Println("Discord message received: ", message.Content)

	if strings.ToLower(message.Content[len(jackal.Discord.CommandPrefix):]) == "ping1" {
		_, err = jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "Pong!")
		jackal.Logger.Info.Println("Received ping, Pong!")
	}

	return
}
