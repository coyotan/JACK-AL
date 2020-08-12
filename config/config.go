package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"../logUtil"
)

var (
	//	logCon	*log.Logger
	logErr       *log.Logger
	possibleCfgs = []string{"./config.json", GetConfDir() + "./config.json"}
)

//Take in a core, which must be compliant with the initInt interface. Let's try to make this lib modular too!
func Init(core initInt) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger) {
	console, info, warn, err = logUtil.InitLoggers(core.LogFile())

	//	logCon = console
	logErr = err

	//Before we load, let's see if it's the first run. If it is, we'll make the config file next.
	if isFirstRun() {

	}

	//Range through possible locations until we find one that exists.
	for _, v := range possibleCfgs {
		if logUtil.VerifyFile(v) {
			//if the file exists, load it.
			LoadCfg(v, &core)
			break
		} else {
			core.LogError().Println("Failed to locate config file.")
			//os.Exit(12)
		}
	}

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

func SaveCfg(core interface{}, fName string) (err error) {
	//Check if fName exists.
	//If it does not exist, make it!
	//If it errors, freak the hell out
	//If it exists, try to open it.
	//If it errors, freak the hell out
	//If it doesn't error, try to write to it.
	//If it errors, freak the hell out.
	return
}

func GetConfDir() (fPath string) {
	path, err := os.UserConfigDir()

	//We can exit code for this, since this shouldn't ever happen.
	if err != nil {
		logErr.Println("Couldn't find config directory.", err)
		//os.Exit(12)
	}

	return path
}

//isFirstRun returns a boolean if the program detects this is its first run. This can be evaluated by checking for the existence of a configuration file.
//If one does not exist at either path, then we can assume that this is the first run.
func isFirstRun() (firstRun bool) {
	for _, v := range possibleCfgs {
		if logUtil.VerifyFile(v) {
			firstRun = false
		} else {
			firstRun = true
		}
	}

	return
}
