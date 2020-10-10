package responders

import (
	"errors"
	"github.com/CoyoTan/JACK-AL/bot/responders/dnd5e"
	"github.com/bwmarrin/discordgo"
	"github.com/txgruppi/parseargs-go"
)

func init() {
	addInitModListener("DND5e", dnd5e.InitDnd)
	addCreateListener("add5eCalendar", addDndGuildCalender)
	addCreateListener("add5ecalendar", addDndGuildCalender)
	addCreateListener("add5eCal", addDndGuildCalender)
	addCreateListener("add5ecal", addDndGuildCalender)
}

func addDndGuildCalender(m *discordgo.Message) (err error) {

	args, err := parseargs.Parse(m.Content)

	if err != nil {
		return
	} else if len(args[1]) < 10 {
		return errors.New("invalid Google Calendar id")
	}

	dnd5e.AddCalBuBuild(m.GuildID, args[1])

	return
}
