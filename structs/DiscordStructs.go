package structs

import (
	"github.com/bwmarrin/discordgo"
)

//DiscordConn contains information which is needed to establish a Discord Session connection, and provide a mount point for command Responders.
type DiscordConn struct {
	User    *discordgo.User    `json:"-"`
	Session *discordgo.Session `json:"-"`

	//Discord Token. This can be saved in json.
	Token string `json:"discordToken"`

	//Discord Maximum State Size. This is how many messages we save from the state pool.
	MaxMessageCount int `json:"maxMessageCount"`

	//Structure information pertaining to dispatching and responding.
	CommandPrefix string `json:"commandPrefix"`

	InitModListeners map[string]func(jackal *CoreCfg) (err error)
	CreateListeners  map[string]func(message *discordgo.Message) (err error)
	DeleteListeners  map[string]func(deleted *discordgo.MessageDelete) (err error)
	EditListeners    map[string]func(deleted *discordgo.MessageUpdate) (err error)
}
