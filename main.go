package main

import (
	"./bot"
	"./config"
	"./structs"
)

var (
	Jackal = structs.CoreCfg{
		LogFilePath: "C:/Users/Coyotan/Documents/JACK-AL/Log.txt",
	}
)

func init () {
	Jackal.Logger.Console, Jackal.Logger.Info, Jackal.Logger.Warn, Jackal.Logger.Error = config.Init(&Jackal)
	Jackal.Logger.Info.Println("Passing to package: Bot")
	bot.Init(&Jackal)
}

func main() {

}

//TODO: SORT OUT EXIT CODES. MAKE SURE THAT THEY ALL LINE UP WITH THE RIGHT ERROR.
//TODO: Create basic Discord connections, and link them into Jackal.
//TODO: Create command Dispatcher and link it into Jackal.
//TODO: Create command Responder template, and link it into Jackal.
//TODO: Project Phase complete. Start testing for responsiveness.