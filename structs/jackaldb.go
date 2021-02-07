package structs

import (
	"github.com/boltdb/bolt"
)

type JackalDB struct {
	db *bolt.DB
}

//InitDB is called by the core loader of the bot each time the program is executed.
func (core *CoreCfg) InitDB() (dbError error) {
	core.DB = &JackalDB{}
	core.DB.db = &bolt.DB{}

	core.DB.db, dbError = bolt.Open(core.GetConfDir()+"/jackal.db", 0640, nil)

	if dbError != nil {
		return dbError

	}

	return
}

//Path returns the location of the db file that was loaded. Wrapping it protects the information.
func (b *JackalDB) Path() string {
	return b.db.Path()
}

//Close completes the database operations and cleanly unloads the database. Wrapping it protects the function.
func (b *JackalDB) Close() {
	b.db.Close()
}
