package main

import (
	"./config"
	"./structs"
)

var (
	Core = structs.CoreCfg{
		LogFilePath: "C:/Users/Coyotan/Documents/JACK-AL/Log.txt",
	}
)

func main() {
	Core.Logger.Console, Core.Logger.Info, Core.Logger.Warn, Core.Logger.Error = config.Init(&Core)
}


//TODO: Create basic Discord connections, and link them into Core.
//TODO: Create command Dispatcher and link it into Core.
//TODO: Create command Responder template, and link it into Core.
//TODO: Project Phase complete. Start testing for responsiveness.