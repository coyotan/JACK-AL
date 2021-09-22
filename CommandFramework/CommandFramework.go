package CommandFramework

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/coyotan/JACK-AL/structs"
)

type CommandFramework struct {
}

type CommandGroup struct {
	Commands map[string]Command
	Name     string
}

//TODO: Database structure (guildID string, channelID string, name string ) PRIMARY KEY, group string, commandPermissions int
//Permissions have hierarchy - Guild permissions apply universally across the server, channel specific permissions can be used to override bot permissions in an isolated channel.
//TODO: Apparently I forgot to add default permissions... well, that's another item.
func (g *CommandGroup) NewCommand(name string, execute interface{}, commContext int, alias ...string) (err error) {

	c := Command{
		Group:      g,
		name:       name,
		Execute:    execute,
		commandCtx: commContext,
		Alts:       alias,
	}

	if g.Commands == nil {
		g.Commands = make(map[string]Command)
	}
	if _, exists := g.Commands[c.name]; exists {
		return errors.New(g.Name + " already contains an entry for " + c.name)
	} else {
		g.Commands[c.Name()] = c
	}

	return nil
}

func (g *CommandGroup) RegisterAllCommands() error {

	for _, c := range g.Commands {
		switch c.Execute.(type) {
		//This section is for bot events, we might even be able to put database return events in here to prevent plugins from directly interacting with the database.
		case func(*discordgo.Session, *structs.CoreCfg) error: //InitEvent
		//TODO: Enumerate other possible JACKAL events which modders might be able to rely on. Consider using events to power database activities, or cross-bot communication?

		//DiscordGo listeners that we restrict the use of, for permissions. Prior to release, this will need to be a full list, so we can remove passing through the entire discord structure.
		case func(*discordgo.Session, *discordgo.MessageCreate) error:
			if _, exists := CreateEventResponders[c.name]; exists {
				return errors.New(c.name + " is already a registered Create event Responder")
			} else {
				CreateEventResponders[c.name] = c
				fmt.Println("Registered new create listener", c.name)
			}

		case func(*discordgo.Session, *discordgo.MessageDelete) error:
			if _, exists := DeleteEventResponders[c.name]; exists {
				return errors.New(c.name + " is already a registered Delete event Responder")
			} else {
				CreateEventResponders[c.name] = c
			}
		default:
			fmt.Println("No match")
		}
	}

	return nil
}

type Command struct {
	Group       *CommandGroup
	Execute     interface{}
	name        string
	Alts        []string
	commandCtx  int //This field uses bitmask permissions to evaluate where the command can be run. Options are: Enabled, DiscordOnly, GuildOnly, GroupOnly, DMOnly, TerminalOnly, APIOnly, and RPCOnly.
	permissions int //This field uses Discord's permissions system. It can be used to establish the minimum permissions required to run the command.

}

func (c *Command) Name() (name string) {
	return c.name
}

func (c *Command) CommandCtx() (value int) {
	return c.commandCtx
}

func (c *Command) Permissions() (value int) {
	return c.permissions
}
