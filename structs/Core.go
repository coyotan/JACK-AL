package structs

import "log"

type CoreCfg struct {
	Token       string `json:"token"`
	Name        string
	Logger      logs
	LogFilePath string `json:"logFile"`

	CommandPrefix string `json:"commandPrefix"`
}

func (core *CoreCfg) LogFile() (fPath string) {
	return core.LogFilePath
}

func (core *CoreCfg) LogConsole() (console *log.Logger) {
	return core.Logger.Console
}

func (core *CoreCfg) LogInfo() (info *log.Logger) {
	return core.Logger.Info
}

func (core *CoreCfg) LogWarn() (warn *log.Logger) {
	return core.Logger.Warn
}

func (core *CoreCfg) LogError() (error *log.Logger) {
	return core.Logger.Error
}

type logs struct {
	Console *log.Logger
	Info    *log.Logger
	Warn    *log.Logger
	Error   *log.Logger
}
