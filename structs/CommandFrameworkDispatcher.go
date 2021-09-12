package structs

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

//TODO: Implement CreateMessage, DeleteMessage, and InitDispatcher events. Translate them from responders framework.
//This should make creating new commands a lot easier for developers looking to add content to Jackal.

var (
	//Jackal Framework Events
	InitEventResponders map[string]Command //TODO: Review this in the future. We probably don't want plugins to have access to CoreCfg.

	//DiscordEvent Handlers
	CreateEventResponders map[string]Command
	DeleteEventResponders map[string]Command
)

const (
	//Jackal Command Context bitmasks.
	Enabled  int = 0x1
	Discord  int = 0x2
	Guild    int = 0x4
	Group    int = 0x8
	DM       int = 0x10
	Terminal int = 0x20
	API      int = 0x40
	RPC      int = 0x80
)

func init() {
	CreateEventResponders = make(map[string]Command)
	DeleteEventResponders = make(map[string]Command)
}

//TODO: Create separate event dispatchers for different inputs, such as terminal, RPC, API, and over Discord. Note, all dispatchers will use the same command map, and check permissions/context to ensure that the command can be run

func AllDispatchInit(session *discordgo.Session, core *CoreCfg) {
	//No context checks. Init event is a broadcast to all listening handlers.
}

//Look at using interface with DiscordGo to register special create handler which will allow us to use this function with all the different types of data.
func DiscordDispatchCreate(session *discordgo.Session, created *discordgo.MessageCreate) {

	//Since this is the Discord command Dispatcher, we should first require that a command be able to be run on Discord, and be Enabled.
	var requiredPermissions = Discord + Enabled

	messageChannel, err := session.Channel(created.ChannelID)

	if err != nil {
		fmt.Println("Failed to find channel")
	}

	fmt.Println("Checking permissions")
	//Based on the channel type, we can assert additional requirements to ensure that the command we're preparing to run can be run in this current context.
	switch messageChannel.Type {
	case discordgo.ChannelTypeGuildText:
		requiredPermissions = requiredPermissions + Guild
	case discordgo.ChannelTypeDM:
		requiredPermissions = requiredPermissions + DM
	case discordgo.ChannelTypeGroupDM:
		requiredPermissions = requiredPermissions + Group
		//TODO: In the future, we should add specific scenarios for News and Store, as these two can have messages sent in them.
	}
	fmt.Println("Required permissions are ", requiredPermissions)

	for _, command := range CreateEventResponders {
		fmt.Println("Present permissions are ", command.commandCtx)
		if (command.commandCtx & requiredPermissions) == requiredPermissions {
			//TODO Check the server to make sure that there are no overrides preventing us from using this command
			c := command.Execute.(func(*discordgo.Session, *discordgo.MessageCreate) error)

			if err := c(session, created); err != nil {
				fmt.Println(err)
			}

		} else {
			//Let the user know that they do not have permission to run that command.
			fmt.Println("Inadequate permissions")
		}
	}
}

/*
func DiscordDispatchDelete(session *discordgo.Session, deleted *discordgo.MessageDelete) {
	//Check context to make sure that the command is allowed to be run in the current context.
}
*/
