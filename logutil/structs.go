package logutil

import (
	"log"
	"os"
)

type UtilLogger struct {
	Level
}

func (u *UtilLogger) PrintFatal(v interface{}, code int) {
	u.Console.Println(v)

	if initComplete {
		u.Error.Println(v)
	}

	os.Exit(code)
}

func (u *UtilLogger) PrintWarn(v interface{}) {
	u.Console.Println(v)
	u.Warning.Println(v)
}

func (u *UtilLogger) PrintInfo(v interface{}) {
	u.Console.Println(v)
	u.Info.Println(v)
}

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
