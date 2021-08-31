package logutil

import (
	"github.com/coyotan/JACK-AL/botutils"
	"log"
	"os"
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
	lFile, err := botutils.CreateFile(botutils.ConfigDir, "/jackal.log")

	if err != nil {
		localLogger.Console.Fatal(err)
		os.Exit(1)
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
