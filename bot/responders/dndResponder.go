package responders

import (
	"github.com/CoyoTan/JACK-AL/bot/responders/dnd5e"
	"github.com/bwmarrin/discordgo"
)

func init() {
	addInitModListener("DND5e", dnd5e.InitDnd)
}

func addDndGuildCalender(m *discordgo.Message) (err error) {

	return
}
