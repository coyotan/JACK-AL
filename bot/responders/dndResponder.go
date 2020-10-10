package responders

import "github.com/CoyoTan/JACK-AL/bot/responders/dnd5e"

func init() {
	addInitModListener("DND5e", dnd5e.InitDnd)
	jackal.Logger.Info.Println("DND5E Initialized!")
}
