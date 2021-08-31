package botutils

import (
	"os"
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
