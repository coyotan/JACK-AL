package main

import (
	"fmt"
	"github.com/coyotan/JACK-AL/bot"
	"github.com/coyotan/JACK-AL/config"
	"github.com/coyotan/JACK-AL/structs"
	"os"
)

var (
	jackal      = structs.CoreCfg{}
	jackalClass = "Esper"
)

func init() {
	jackal.Logger.Console, jackal.Logger.Info, jackal.Logger.Warn, jackal.Logger.Error = config.Init(&jackal)
	jackal.Logger.Info.Println("\n\n//========== JACK-AL: " + jackalClass + " Has Begun Execution. ==========\\\\")
	jackal.Logger.Info.Println("Passing to package: Bot")

	//This is here to support containers, because APPARENTLY making life easy requires first making it more difficult.
	if config.IsDockerContainer() {
		jackal.Discord.Token = os.Getenv("DISCTOKEN")
		if len(os.Getenv("DISCTOKEN")) > 1 {
			jackal.Logger.Info.Println("Grabbed DISCTOKEN from environment")
		} else {
			jackal.Logger.Error.Println("DISCTOKEN is not present in the environment. This WILL cause the container to exit on code 101 if this is the first run, or a volume has not been configured. Please specify the token in the environment.")
			fmt.Println("DISCTOKEN not present in environment. Please read the logs for more information.")
		}
	}

	bot.Init(&jackal)
}

func main() {

}
