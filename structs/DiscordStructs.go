package structs

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordConn struct {
	User		*discordgo.User		`json:"-"`
	Session		*discordgo.Session	`json:"-"`

	//Discord Token. This can be saved in json.
	Token		string `json:discordToken`

	//Structure information pertaining to dispatching and responding.
	CommandPrefix	string 		`json:"commandPrefix"`

	CreateListeners map[string]func(message *discordgo.Message)(err error)
}
