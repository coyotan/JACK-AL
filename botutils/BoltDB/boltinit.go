package BoltDB

import (
	"github.com/CoyoTan/JACK-AL/structs"
	"github.com/boltdb/bolt"
	"os"
)

var (
	jackal *structs.CoreCfg
)

//InitBoltDB is called during the bot configuration loading process.
func InitBoltDB(jack *structs.CoreCfg) {
	jackal = jack
	path := getConfDir()

	jackal.DB, err := bolt.Open(path+"jackal.db", 0600, nil)
	if err != nil {

	}

}

//GetConfDir returns the file location for the ideal place to use for a working directory.
func getConfDir() (fPath string) {
	path, err := os.UserConfigDir()

	//We can exit code for this, since this shouldn't ever happen.
	if err != nil {
		jackal.Logger.Error.Println("Couldn't find config directory.", err)
		os.Exit(12)
	}

	return path + "/JACK-AL"
}
