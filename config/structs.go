package config

import (
	"log"
)

type initInt interface {
	LogFile() string

	LogConsole() *log.Logger
	LogInfo() *log.Logger
	LogWarn() *log.Logger
	LogError() *log.Logger
	GetConfDir() string
}
