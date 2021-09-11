package config

import (
	"fmt"
	"github.com/coyotan/JACK-AL/botutils"
	"github.com/coyotan/JACK-AL/logutil"
	"log"
)

var (
	logErr     *log.Logger
	ConfigPath string
)

//Init takes in a core, which must be compliant with the initInt interface. Let's try to make this lib modular too!
func Init(core initInt) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger) {
	console, info, warn, err = logutil.InitLoggers()

	logErr = err

	ConfigPath = botutils.ConfigDir + "/config.json"

	//Before we load, let's see if it's the first run. If it is, we'll make the config file next.
	if botutils.IsFirstRun() {

		if !botutils.IsDockerContainer() {
			fmt.Println("JACK-AL has detected that this is the first time it's been here. Please go to " + botutils.ConfigDir + "and populate the config.json file with the bot's Discord Token.")
		}

		//Make file in ConfDir, and return as file used, so we can adjust the code that follows...
		err := botutils.SaveCfg(ConfigPath, &core)

		if err != nil {
			logErr.Println("There was a critical error loading ")
		}
	}

	_ = botutils.LoadCfg(ConfigPath, &core)

	return
}
