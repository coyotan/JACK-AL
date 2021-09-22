package CommandFramework

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/coyotan/JACK-AL/structs"
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
	//Jackal Command Constants for Context and Permission bitmask.
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

func AllDispatchInit(session *discordgo.Session, core *structs.CoreCfg) {
	//No context checks. Init event is a broadcast to all listening handlers.
}

//Look at using interface with DiscordGo to register special create handler which will allow us to use this function with all the different types of data.
func DiscordDispatchCreate(session *discordgo.Session, created *discordgo.MessageCreate) {

	//Since this is the Discord command Dispatcher, these are the minimum requirements.
	var requiredContext = Discord + Enabled

	messageChannel, err := session.Channel(created.ChannelID)

	if err != nil {
		fmt.Println("Failed to find channel")
	}

	//Based on the channel type, we can assert additional requirements to ensure that the command we're preparing to run can be run in this current context.
	switch messageChannel.Type {
	case discordgo.ChannelTypeGuildText:
		requiredContext = requiredContext + Guild
	case discordgo.ChannelTypeDM:
		requiredContext = requiredContext + DM
	case discordgo.ChannelTypeGroupDM:
		requiredContext = requiredContext + Group
		//TODO: In the future, we should add specific scenarios for News and Store, as these two can have messages sent in them.
	}

	for _, command := range CreateEventResponders {
		fmt.Println("Present context is ", command.commandCtx)
		if (command.commandCtx & requiredContext) == requiredContext {
			//TODO Query the User/GuildCommands table for permission. Check for a channel entry, if no channel entry, check for a guild entry, if no guild entry, assume default required permissions. AND user permissions against requiredPermissions.
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
