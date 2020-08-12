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
