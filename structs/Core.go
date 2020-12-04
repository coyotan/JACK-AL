package structs

import (
	"log"
)

//CoreCfg is the most important part of this software, contains all information used by every part of this program. This is the heart of JACK-AL
type CoreCfg struct {
	//Structure components pertaining specifically to logging.
	Logger      logs
	LogFilePath string `json:"logFile"`

	DB *JackalDB //We do not need any DB functions in our interface because we do not need this to be saved. This is a note for future me.

	Discord DiscordConn `json:"discord"`
}

//LogFile is an exported function which returns the location of the LogFile. This is relevant to create logging interfaces.
func (core *CoreCfg) LogFile() (fPath string) {
	return core.LogFilePath
}

//LogConsole is an exported function which returns a pointer to the console logger, which can be used to share informational messages with a console user.
func (core *CoreCfg) LogConsole() (console *log.Logger) {
	return core.Logger.Console
}

//LogInfo is an exported function which returns a pointer to the Info logger, which can be used to save informational messages to a log file.
func (core *CoreCfg) LogInfo() (info *log.Logger) {
	return core.Logger.Info
}

//LogWarn is an exported function which returns a pointer to the Warning logger, which can be used to save potentially disruptive  messages to a log file.
func (core *CoreCfg) LogWarn() (warn *log.Logger) {
	return core.Logger.Warn
}

//LogError is an exported function which returns a pointer to the Error logger, which can be used to save Emergency/Dangerous messages to a log file.
func (core *CoreCfg) LogError() (error *log.Logger) {
	return core.Logger.Error
}

type logs struct {
	Console *log.Logger
	Info    *log.Logger
	Warn    *log.Logger
	Error   *log.Logger
}
