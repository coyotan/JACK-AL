package logutil

import (
	"github.com/coyotan/JACK-AL/botutils"
	"log"
	"os"
	"path/filepath"
)

var (
	initComplete = false
	lFile        *os.File

	//While it is a small waste, by having a copy of the Core Logger here, we can make this software modular and avoid cycling imports.
	localLogger UtilLogger
)

//Initialize the library. Right now, we don't really need to do anything here.
func init() {
}

//InitLoggers will initialize and return all the log handlers. We're going to try to do this modularly.
func InitLoggers() (Console *log.Logger, Info *log.Logger, Warn *log.Logger, Error *log.Logger) {

	//Init console logger first. We can know for sure that this one is going to work.
	Console = log.New(os.Stdout, "Console: ", log.Ltime|log.Lshortfile)
	localLogger.Console = Console

	//Create the file to start overwrite ... without this, we're getting some stupid bug.
	lFile, _ = CreateFile(botutils.ConfigDir, "/jackal.log")

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

//CreateFile will attempt to create a file, and if file creation for the log file fails, flip shit.
func CreateFile(path string, fName string) (file *os.File, err error) {

	err = os.MkdirAll(path, 0644)

	if err != nil {
		if initComplete {
			localLogger.Error.Println("Path " + path + " could not be created!\n" + err.Error())
		}
	}

	file, err = os.OpenFile(filepath.Clean(filepath.Join(path, fName)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	if err != nil {
		if initComplete {
			//Log error creating file
			localLogger.Error.Println("File " + fName + " could not be created!\n" + err.Error())
		}
		//This is sloppy and missing a check, but the log file SHOULD be the only file we attempt to create before logging is enabled.
		localLogger.Console.Println("(╯°□°）╯︵ ┻━┻")
		localLogger.PrintFatal("A critical error prevented the creation of the log file. Execution will not continue.\n"+err.Error(), 1)
		//Exit code 1 is reserved for failed creation of log file. This should be a dead give away of the issue.
	}

	return file, err
}

//We will need to add support for hunting down filepaths and finding the folder that does not exist. Program does not automatically identify that directories need to be made.
