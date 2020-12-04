package structs

import "github.com/boltdb/bolt"

type JackalDB struct {
	db *bolt.DB
}

//Wrapper function which prevents this from being written to.
func (b *JackalDB) Path() string {
	return b.db.Path()
}

//Wrapper function which prevents this from being written to.
func (b *JackalDB) Close() {
	b.db.Close()
}
