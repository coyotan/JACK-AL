package botutils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

//LoadCfg configuration from specified directory and provide it to interface. Ideally this would be a reference.
func LoadCfg(filename string, config interface{}) (err error) {

	file, err := os.Open(filepath.Clean(filename))

	if err != nil {
		return err
		//Unreachable but left for future documentation
		os.Exit(10)
	}

	defer file.Close()

	byteVal, err := ioutil.ReadAll(file)
	if err != nil {
		return err
		//Unreachable but left for future documentation
		os.Exit(11)
	}

	err = json.Unmarshal(byteVal, &config)

	return err
}

//SaveCfg will save running configurations to a provided file, or respond with an error if something goes wrong.
func SaveCfg(fName string, core interface{}) (err error) {

	if VerifyFile(fName) {
		confOut, err := json.Marshal(core)
		if err != nil {
			return err
			//Unreachable but left for future documentation
			os.Exit(14)
			//File failed to write to config... but it did open.
		}

		//TODO: At a later date, change this from a discard to some logging information for posterity's sake.
		_ = ioutil.WriteFile(fName, confOut, 600)

	} else {
		//If it does not exist, make it!
		err = os.MkdirAll(fName[:len(fName)-12], 750)

		if err != nil {
			return err
			//Unreachable but left for future documentation
			os.Exit(13)
		}

		_, err = os.Create(fName)
		if err != nil {
			return err
			//Unreachable but left for future documentation
			os.Exit(12)
			//Prevent recursion by stopping here... we do NOT want to continue with this one.
		}

		_ = SaveCfg(fName, core)
	}

	return
}
