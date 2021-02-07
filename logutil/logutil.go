package logutil

import (
	"log"
	"os"
)

var (
	initComplete = false
	lFile        *os.File

	//While it is a small waste, by having a copy of the Core Logger here, we can make this software modular and avoid cycling imports.
	localLogger Level
)

//Initialize the library. Right now, we don't really need to do anything here.
func init() {
}

//InitLoggers will initialize and return all the log handlers. We're going to try to do this modularly.
func InitLoggers(logFile string) (Console *log.Logger, Info *log.Logger, Warn *log.Logger, Error *log.Logger) {

	//Init console logger first. We can know for sure that this one is going to work.
	Console = log.New(os.Stdout, "Console: ", log.Ltime|log.Lshortfile)
	localLogger.Console = Console

	//Create the file to start overwrite ... without this, we're getting some stupid bug.
	if len(logFile) < 4 {
		localLogger.Info.Println("Choosing default logging location.")
		localLogger.Warning.Println("If you have specified a new log location, make sure that the file path contains more than 4 characters.")
		lFile, _ = CreateFile(GetUserConfDir() + "/jackal.log")
	} else {
		localLogger.Info.Println("Choosing specified logging location.")
		lFile, _ = CreateFile(logFile)
	}

	//Now that we know the file exists, we can use the rest of these.
	Info = log.New(lFile, "INFO: ", log.Ltime|log.Ldate|log.Lshortfile)
	Warn = log.New(lFile, "WARN: ", log.Ltime|log.Ldate|log.Lshortfile)
	Error = log.New(lFile, "ERROR: ", log.Ltime|log.Ldate|log.Lshortfile)

	//Add some stuff for us to take care of ourselves in here. In the name of modularity.
	localLogger.Info = Info
	localLogger.Warning = Warn
	localLogger.Error = Error

	//Tell everything we're all taken care of!
	initComplete = true
	Info.Println("Logging initialized.")

	return
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

//CreateFile will attempt to create a file, and if file creation for the log file fails, flip shit.
func CreateFile(fName string) (fHandle *os.File, err error) {

	fHandle, err = os.OpenFile(fName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		if initComplete {
			//Log error creating file
			localLogger.Error.Println("File " + fName + " could not be created!\n" + err.Error())
		}
		//This is sloppy and missing a check, but the log file SHOULD be the only file we attempt to create before logging is enabled.
		localLogger.Console.Println("(╯°□°）╯︵ ┻━┻")
		localLogger.Fatal("A critical error prevented the creation of the log file. Execution will not continue.\n"+err.Error(), 1)
		//Exit code 1 is reserved for failed creation of log file. This should be a dead give away of the issue.
	}

	return fHandle, err
}

func GetUserConfDir() (path string) {

	if path, err := os.UserConfigDir(); err != nil {
		localLogger.Fatal("There was a critical error attempting to access the logging directory\n"+err.Error(), 3)
		return ""
	} else {
		return path + "/JACK-AL"
	}
}

//We will need to add support for hunting down filepaths and finding the folder that does not exist. Program does not automatically identify that directories need to be made.
