package main

import (
	"github.com/CoyoTan/JACK-AL/bot"
	"github.com/CoyoTan/JACK-AL/config"
	"github.com/CoyoTan/JACK-AL/structs"
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
