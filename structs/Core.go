package structs

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"os"
)

//CoreCfg is the most important part of this software, contains all information used by every part of this program. This is the heart of JACK-AL
type CoreCfg struct {
	//Structure components pertaining specifically to logging.
	Logger   logs
	Database Db
	Discord  DiscordConn `json:"discord"`
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

//GetConfDir returns the file location for the ideal place to use for a working directory.
func (core *CoreCfg) GetConfDir() (fPath string) {
	path, err := os.UserConfigDir()

	//We can exit code for this, since this shouldn't ever happen.
	if err != nil {
		core.Logger.Error.Println("Couldn't find config directory.", err)
		os.Exit(12)
	}

	return path + "/JACK-AL"
}

//VerifyFile returns false if the filename present does not exist in the filesystem.
//Exported because it makes writing other things that need to use this a lot smoother.
func (core *CoreCfg) VerifyFile(fName string) (fExists bool) {

	if _, err := os.Stat(fName); os.IsNotExist(err) {
		fExists = false
	} else {
		fExists = true
	}

	return fExists
}

func (core *CoreCfg) InitCassandraDB() (err error) {

	//This line will need to be changed when we add support for multiple clusters in the end.
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA"))
	//TODO We need to find a way to check if this exists before we use it to connect. This will involve more first time checking.
	cluster.Keyspace = "jackal"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()

	core.Database.session = session

	if err != nil {
		core.Logger.Error.Println("could not connect to Cassandra cluster ", err)
		fmt.Println("There was an error connecting to the Cassandra cluster. Please check the logs for more information.")
		return err
	}

	if err := session.Query("SELECT * FROM jackal.users;").Exec(); err != nil {
		return err
	}

	if err := core.Database.CreateKeyspace(); err != nil {
		return err
	}

	if err = core.Database.CreateUserTable(); err != nil {
		return err
	}

	if err = core.Database.CreateMessagesTable(); err != nil {
		return err
	}

	return err
}
