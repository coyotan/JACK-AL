package main

import (
	"github.com/CoyoTan/JACK-AL/bot"
	"github.com/CoyoTan/JACK-AL/config"
	"github.com/CoyoTan/JACK-AL/structs"
)

var (
	jackal = structs.CoreCfg{
		LogFileDir:  "F:/Documents/JACK-AL",
		LogFilePath: "F:/Documents/JACK-AL/Log.txt",
	}
)

func init() {
	jackal.Logger.Console, jackal.Logger.Info, jackal.Logger.Warn, jackal.Logger.Error = config.Init(&jackal)
	jackal.Logger.Info.Println("Passing to package: Bot")
	bot.Init(&jackal)
}

func main() {

}
