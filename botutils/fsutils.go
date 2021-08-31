package botutils

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ConfigDir, _ = getConfDir()
	possibleCfgs = []string{"./config.json", ConfigDir + "/config.json"}
)

//GetUserConfDir gets the application data directory of the operating system this code is running on. For example, in Windows this is %APPDATA%/JACK-AL
func getConfDir() (path string, err error) {
	path, err = os.UserConfigDir()

	if err != nil {
		return "", err
	}
	return path + "/JACK-AL", err
}

//VerifyFile returns false if the filename present does not exist in the filesystem.
//Exported because it makes writing other things that need to use this a lot smoother.
func VerifyFile(fName string) (fExists bool) {

	if _, err := os.Stat(fName); os.IsNotExist(err) {
		fExists = false
	} else {
		fExists = true
	}

	return fExists
}

//IsDockerContainer checks to evaluate if the bot is running in a docker container, or on bare metal.
func IsDockerContainer() (IsContainer bool) {
	if VerifyFile("/.dockerenv") {
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
		if VerifyFile(v) {
			firstRun = false
			break
		} else {
			firstRun = true
		}
	}
	return
}

//CreateFile will attempt to create a file, and if file creation for the log file fails, flip shit.
func CreateFile(path string, fName string) (file *os.File, err error) {

	err = os.MkdirAll(path, 0644)

	if err != nil {
		return nil, errors.New("path " + path + " could not be created\n" + err.Error())
	}

	file, err = os.OpenFile(filepath.Clean(filepath.Join(path, fName)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	if err != nil {
		return nil, errors.New("(╯°□°）╯︵ ┻━┻\n A critical error prevented the creation of " + fName + "\n" + err.Error())
	}

	return file, err
}

//We will need to add support for hunting down filepaths and finding the folder that does not exist. Program does not automatically identify that directories need to be made.
