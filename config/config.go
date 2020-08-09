package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"../logUtil"
)

var (
	core initInt
)

//Take in a core, which must be compliant with the initInt interface. Let's try to make this lib modular too!
func Init(core initInt) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger) {
	console, info, warn, err = logUtil.InitLoggers(core.LogFile())
	return
}

//Load configuration from specified directory and provide it to interface. Ideally this would be a reference.
func LoadCfg(filename string, config interface{}) (err error) {

	logError := core.LogError()

	file, err := os.Open(filename)

	if err != nil {
		logError.Println("There was a critical error opening the core json.\n", err)
		os.Exit(10)
	}

	defer file.Close()

	byteVal, err := ioutil.ReadAll(file)
	if err != nil {
		logError.Println("There was a critical error reading data from the config file.\n", err)
		os.Exit(11)
	}

	json.Unmarshal(byteVal, &config)

	return err
}
