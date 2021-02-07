package logutil

import (
	"log"
	"os"
)

//UtilLogger is a wrapper struct to allow for the injection of functions we can use to increase verbosity. These will be updated later with a debug option for the CLI.
type UtilLogger struct {
	Level
}

//PrintFatal is a wrapper function which prints Fatal logs to the console and to the log file, for increased verbosity.
func (u *UtilLogger) PrintFatal(v interface{}, code int) {
	u.Console.Println(v)

	if initComplete {
		u.Error.Println(v)
	}

	os.Exit(code)
}

//PrintWarn is a wrapper function which prints Warning logs to the console and to the log file, for increased verbosity.
func (u *UtilLogger) PrintWarn(v interface{}) {
	u.Console.Println(v)
	u.Warning.Println(v)
}

//PrintInfo is a wrapper function which prints Informative logs to the console and to the log file, for increased verbosity.
func (u *UtilLogger) PrintInfo(v interface{}) {
	u.Console.Println(v)
	u.Info.Println(v)
}

//PrintConsole is a wrapper function for ease of use.
func (u *UtilLogger) PrintConsole(v interface{}) {
	u.Console.Println(v)
}

//Level contains different loggers to provide alternative logging prefixes and configurations, to provide relevant information and highlight relevant events.
type Level struct {
	Console *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}
