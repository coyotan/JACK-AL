package structs

import "log"

type CoreCfg struct {
	//Structure components pertaining specifically to logging.
	Logger      	logs
	LogFilePath 	string 		`json:"logFile"`

	Discord			DiscordConn `json:"discord"`
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
