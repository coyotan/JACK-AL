package logUtil

import (
	"log"
	"os"
)

type Level struct {
	Console *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

//Our re-implementation of fatal logger. Allows for custom generated exit codes, which we can use to better diagnose failure.
func (l *Level) Fatal(v interface{}, code int) {
	l.Console.Println(v)

	//Gotta make sure that we've initialized the rest of the logging system.
	if initComplete {
		l.Error.Println(v)
	}

	os.Exit(code)
}