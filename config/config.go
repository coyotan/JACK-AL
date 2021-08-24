package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/coyotan/JACK-AL/logutil"
)

var (
	logErr       *log.Logger
	possibleCfgs = []string{"./config.json", GetConfDir() + "/config.json"}
	configPath   string
)

//Init takes in a core, which must be compliant with the initInt interface. Let's try to make this lib modular too!
func Init(core initInt) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger) {
	console, info, warn, err = logutil.InitLoggers()

	logErr = err

	fPath := GetConfDir()
	configPath = fPath + "/config.json"

	//Before we load, let's see if it's the first run. If it is, we'll make the config file next.
	if IsFirstRun() {

		if !IsDockerContainer() {
			fmt.Println("JACK-AL has detected that this is the first time it's been here. Please go to " + GetConfDir() + "and populate the config.json file with the bot's Discord Token.")
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

//SaveCfg will save running configurations to a provided file, or respond with an error if something goes wrong.
func SaveCfg(fName string, core interface{}) (err error) {

	if logutil.VerifyFile(fName) {
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

//GetConfDir returns the file location for the ideal place to use for a working directory.
func GetConfDir() (fPath string) {

	path, err := os.UserConfigDir()
	//We can exit code for this, since this shouldn't ever happen.
	if err != nil {
		logErr.Println("Couldn't find config directory.", err)
		os.Exit(12)
	}

	return path + "/JACK-AL"
}

//IsDockerContainer checks to evaluate if the bot is running in a docker container, or on bare metal.
func IsDockerContainer() (IsContainer bool) {
	if logutil.VerifyFile("/.dockerenv") {
		return true
	} else {
		return false
	}
}

//IsFirstRun returns a boolean if the program detects this is its first run. This can be evaluated by checking for the existence of a configuration file.
//If one does not exist at either path, then we can assume that this is the first run.
//TODO: Add support for accessing environment variables to perform authentication to Discord and the Cassandra database.
func IsFirstRun() (firstRun bool) {

	for _, v := range append(possibleCfgs) {
		if logutil.VerifyFile(v) {
			firstRun = false
			break
		} else {
			firstRun = true
		}
	}
	return
}
