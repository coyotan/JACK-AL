package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"../logUtil"
)

var (
	logErr       *log.Logger
	possibleCfgs = []string{"./config.json", GetConfDir() + "/config.json"}
	configPath   string
)

//Take in a core, which must be compliant with the initInt interface. Let's try to make this lib modular too!
func Init(core initInt) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger) {
	console, info, warn, err = logUtil.InitLoggers(core.LogFile())

	logErr = err

	fPath := GetConfDir()
	configPath = fPath + "/config.json"

	//Before we load, let's see if it's the first run. If it is, we'll make the config file next.
	if isFirstRun() {
		//Make file in ConfDir, and return as file used, so we can adjust the code that follows...
		err := SaveCfg(configPath, &core)

		if err != nil {
			logErr.Println("There was a critical error loading ")
		}
	}

	LoadCfg(configPath, &core)

	return
}

//Load configuration from specified directory and provide it to interface. Ideally this would be a reference.
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

//FIXME Be aware, review for possible recursion issues.
func SaveCfg(fName string, core interface{}) (err error) {

	if logUtil.VerifyFile(fName) {
		confOut, err := json.Marshal(core)
		if err != nil {
			logErr.Println("There was a critical error writing save data to the configuration file.\n", err)
			os.Exit(14)
			//File failed to write to config... but it did open.
		}

		err = ioutil.WriteFile(fName, confOut, 600)

	} else {
		//If it does not exist, make it!
		err = os.MkdirAll(fName[:len(fName)-12], 600)

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

func GetConfDir() (fPath string) {
	path, err := os.UserConfigDir()

	//We can exit code for this, since this shouldn't ever happen.
	if err != nil {
		logErr.Println("Couldn't find config directory.", err)
		os.Exit(12)
	}

	return path + "/JACK-AL"
}

//isFirstRun returns a boolean if the program detects this is its first run. This can be evaluated by checking for the existence of a configuration file.
//If one does not exist at either path, then we can assume that this is the first run.
func isFirstRun() (firstRun bool) {

	for _, v := range possibleCfgs {
		if logUtil.VerifyFile(v) {
			firstRun = false
			break
		} else {
			firstRun = true
		}
	}

	return
}
