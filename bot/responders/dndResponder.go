package responders

import "github.com/CoyoTan/JACK-AL/bot/responders/dnd5e"

func init() {
	dnd5e.InitDnd(jackal)
	jackal.Logger.Info.Println("DND5E Initialized!")
}
