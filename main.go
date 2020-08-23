package main

import (
	"./bot"
	"./config"
	"./structs"
)

var (
	jackal = structs.CoreCfg{
		LogFilePath: "C:/Users/Coyotan/Documents/JACK-AL/Log.txt",
	}
)

func init() {
	jackal.Logger.Console, jackal.Logger.Info, jackal.Logger.Warn, jackal.Logger.Error = config.Init(&jackal)
	jackal.Logger.Info.Println("Passing to package: Bot")
	bot.Init(&jackal)
}

func main() {

}

//TODO: SORT OUT EXIT CODES. MAKE SURE THAT THEY ALL LINE UP WITH THE RIGHT ERROR.
//TODO: Project Phase complete. Start testing for responsiveness.
