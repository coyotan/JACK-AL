package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/CoyoTan/JACK-AL/logutil"
)

var (
	logErr       *log.Logger
	possibleCfgs = []string{"./config.json", GetConfDir() + "/config.json"}
	configPath   string
)

//Init takes in a core, which must be compliant with the initInt interface. Let's try to make this lib modular too!
func Init(core initInt) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger) {
	console, info, warn, err = logutil.InitLoggers(core.LogFile())

	logErr = err

	fPath := GetConfDir()
	configPath = fPath + "/config.json"

	//Before we load, let's see if it's the first run. If it is, we'll make the config file next.
	if IsFirstRun() {
		//Make file in ConfDir, and return as file used, so we can adjust the code that follows...
		err := SaveCfg(configPath, &core)

		if err != nil {
			logErr.Println("There was a critical error loading ")
		}
	}

	LoadCfg(configPath, &core)

	return
}

//LoadCfg configuration from specified directory and provide it to interface. Ideally this would be a reference.
func LoadCfg(filename string, config interface{}) (err error) {

	file, err := os.Open(filename)

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

	json.Unmarshal(byteVal, &config)

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

		ioutil.WriteFile(fName, confOut, 660)

	} else {
		//If it does not exist, make it!
		err = os.MkdirAll(fName[:len(fName)-12], 660)

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

		SaveCfg(fName, core)
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

//IsFirstRun returns a boolean if the program detects this is its first run. This can be evaluated by checking for the existence of a configuration file.
//If one does not exist at either path, then we can assume that this is the first run.
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
