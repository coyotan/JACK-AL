package config

import (
	"encoding/json"
	"fmt"
	"github.com/coyotan/JACK-AL/botutils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/coyotan/JACK-AL/logutil"
)

var (
	logErr     *log.Logger
	configPath string
)

//Init takes in a core, which must be compliant with the initInt interface. Let's try to make this lib modular too!
func Init(core initInt) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger) {
	console, info, warn, err = logutil.InitLoggers()

	logErr = err

	configPath = botutils.ConfigDir + "/config.json"

	//Before we load, let's see if it's the first run. If it is, we'll make the config file next.
	if botutils.IsFirstRun() {

		if !botutils.IsDockerContainer() {
			fmt.Println("JACK-AL has detected that this is the first time it's been here. Please go to " + botutils.ConfigDir + "and populate the config.json file with the bot's Discord Token.")
		}

		//Make file in ConfDir, and return as file used, so we can adjust the code that follows...
		err := SaveCfg(configPath, &core)

		if err != nil {
			logErr.Println("There was a critical error loading ")
		}
	}

	_ = LoadCfg(configPath, &core)

	return
}

//TODO: Refactor and remove this. It was added into botutils/fsutils.
//LoadCfg configuration from specified directory and provide it to interface. Ideally this would be a reference.
func LoadCfg(filename string, config interface{}) (err error) {

	file, err := os.Open(filepath.Clean(filename))

	if err != nil {
		logErr.Println("There was a critical error opening the core json.\n", err)
		os.Exit(10)
	}

	defer file.Close()

	byteVal, err := ioutil.ReadAll(file)
	if err != nil {
		logErr.Println("There was a critical error reading data from the config file.\n", err)
		os.Exit(11)
	}

	err = json.Unmarshal(byteVal, &config)

	return err
}

//TODO: Refactor and remove this. It was added into botutils/fsutils.
//SaveCfg will save running configurations to a provided file, or respond with an error if something goes wrong.
func SaveCfg(fName string, core interface{}) (err error) {

	if botutils.VerifyFile(fName) {
		confOut, err := json.Marshal(core)
		if err != nil {
			logErr.Println("There was a critical error writing save data to the configuration file.\n", err)
			os.Exit(14)
			//File failed to write to config... but it did open.
		}

		//TODO: At a later date, change this from a discard to some logging information for posterity's sake.
		_ = ioutil.WriteFile(fName, confOut, 600)

	} else {
		//If it does not exist, make it!
		err = os.MkdirAll(fName[:len(fName)-12], 750)

		if err != nil {
			logErr.Println("There was a critical error creating a directory in "+fName[:12], err)
			os.Exit(13)
		}

		_, err = os.Create(fName)
		if err != nil {
			logErr.Println("There was a critical error creating the save file.", err)
			os.Exit(12)
			//Prevent recursion by stopping here... we do NOT want to continue with this one.
		}

		_ = SaveCfg(fName, core)
	}

	return
}
