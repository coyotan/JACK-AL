package main

import (
	"fmt"
	"github.com/coyotan/JACK-AL/bot"
	"github.com/coyotan/JACK-AL/botutils"
	"github.com/coyotan/JACK-AL/config"
	"github.com/coyotan/JACK-AL/structs"
	"os"
)

var (
	jackal      = structs.CoreCfg{}
	jackalClass = "Esper"
)

func init() {
	if botutils.IsFirstRun() {
		if !botutils.IsDockerContainer() {
			fmt.Println(
				"Hello! My name is JACK-AL, but some people chose to call me Jack, or Jackie. I prefer Jackal, but you're welcome to call me whatever you'd like.\n\nI was created by the Discord user Coyotan#0962. " +
					"His name is Wesley. He told me how much time and effort it took to build me, and that for over three years, I was a constantly developing project. " +
					"When he first woke me up, he told me that even though I was ready now, he would continue to help me grow." +
					"He told me that whenever I meet new people, I should always do my best to help them out, so I've created a basic configuration template for you." +
					"\n\n" +
					"Please go to " + botutils.ConfigDir + " and fill out the configuration template. Once you're done, come back and we can get started!" +
					"\n\nI'm so happy to meet a new friend.")
			config.Init(&jackal)
			os.Exit(0)
		} else {
			fmt.Println("Hello! Please check out " + botutils.ConfigDir + " for optional configuration options, if you are not using environment variables. Note, of my special skills may still require you to make changes to the config file stored in that location.")
		}
	}

	jackal.Logger.Console, jackal.Logger.Info, jackal.Logger.Warn, jackal.Logger.Error = config.Init(&jackal)
	jackal.Logger.Info.Println("\n\n//========== JACK-AL: " + jackalClass + " Has Begun Execution. ==========\\\\")
	jackal.Logger.Info.Println("Passing to package: Bot")

	//This is here to support containers, because APPARENTLY making life easy requires first making it more difficult.
	if botutils.IsDockerContainer() {
		jackal.Discord.Token = os.Getenv("DISCTOKEN")
		//TODO: Add CommandPrefix environment variable as well.
		if len(os.Getenv("DISCTOKEN")) > 1 {
			fmt.Println("Retrieved Bot Token")
		} else {
			fmt.Println("DISCTOKEN is not present in the environment. This WILL cause the container to exit on code 101 if this is the first run, or a volume has not been configured. Please specify the token in the environment.")
			fmt.Println("DISCTOKEN not present in environment. Please read the logs for more information.")
		}
	}

	bot.Init(&jackal)
}

func main() {

}
