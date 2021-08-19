package main

import (
	"github.com/coyotan/JACK-AL/bot"
	"github.com/coyotan/JACK-AL/config"
	"github.com/coyotan/JACK-AL/structs"
)

var (
	jackal      = structs.CoreCfg{}
	jackalClass = "Esper"
)

func init() {
	jackal.Logger.Console, jackal.Logger.Info, jackal.Logger.Warn, jackal.Logger.Error = config.Init(&jackal)
	jackal.Logger.Info.Println("\n\n//========== JACK-AL: " + jackalClass + " Has Begun Execution. ==========\\\\")
	jackal.Logger.Info.Println("Passing to package: Bot")
	bot.Init(&jackal)
}

func main() {

}
