package logUtil

import (
	"log"
	"os"

	"../config"
)

var (
	InitComplete = false
	lFile        *os.File
)

//Initialize the library AND add the logger to the Core structure. This will make it accessible everywhere.
func Init() {

	//Init console logger first. We can know for sure that this one is going to work.
	config.Core.Logger.Console = log.New(os.Stdout, "Console: ", log.Ltime|log.Lshortfile)

	//Just to prevent stupid mistakes.
	if !VerifyFile(config.Core.LogFile) {
		lFile, _ = CreateFile(config.Core.LogFile)
	}

	//Now that we know the file exists, we can use the rest of these.
	config.Core.Logger.Info = log.New(lFile, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	config.Core.Logger.Warning = log.New(lFile, "WARN: ", log.Ltime|log.Ldate|log.Lshortfile)
	config.Core.Logger.Error = log.New(lFile, "", log.Ltime|log.Ldate|log.Lshortfile)

	//Tell everything we're all taken care of!
	InitComplete = true
	config.Core.Logger.Info.Println("Logging initialized.")
}

//VerifyFile returns false if the filename present does not exist in the filesystem.
func VerifyFile(fName string) (fExists bool) {

	if _, err := os.Stat(fName); os.IsNotExist(err) {
		fExists = false
	} else {
		fExists = true
	}

	return fExists
}

//Attempt to create a file, and if file creation for the log file fails, flip shit.
func CreateFile(fName string) (fHandle *os.File, err error) {

	fHandle, err = os.Create(fName)

	if err != nil {
		if InitComplete {
			//Log error creating file
			config.Core.Logger.Error.Println("File " + fName + " could not be created!\n" + err.Error())
		} else {
			//This is sloppy and missing a check, but the log file SHOULD be the only file we attempt to create before logging is enabled.
			config.Core.Logger.Error.Println("(╯°□°）╯︵ ┻━┻")
			config.Core.Logger.Fatal("A critical error prevented the creation of the log file. Execution will not continue.", 1)
			//Exit code 1 is reserved for failed creation of log file. This should be a dead give away of the issue.
		}
	}

	return fHandle, err
}
