package structs

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

//JackalDB contains the boltDatabase. This is used to protect the data in the boltDatabase by providing proxied access to the records.
type JackalDB struct {
	Db *bolt.DB
}

//InitDB is called by the core loader of the bot each time the program is executed.
func (core *CoreCfg) InitDB() (dbError error) {
	core.DB = &JackalDB{}
	core.DB.Db = &bolt.DB{}

	core.DB.Db, dbError = bolt.Open(core.GetConfDir()+"/jackal.Db", 0640, &bolt.Options{Timeout: 5 * time.Second})

	if dbError != nil {
		return dbError
	}
	return
}

//Path returns the location of the Db file that was loaded. Wrapping it protects the information.
func (b *JackalDB) Path() string {
	return b.Db.Path()
}

//Close completes the database operations and cleanly unloads the database. Wrapping it protects the function.
func (b *JackalDB) Close() {
	b.Db.Close()
}

///TODO: Review this function. Wrote this during session.

//Put is a wrapping function for putting data into the databucket database. It should make storage actions simple and safe.
func (b *JackalDB) Put(bucket string, key string, value string) (err error) {

	err = b.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(value))
	})

	//Only if an error has not already been thrown, qErr has an error, or answer the answer is not identical to the value that should have been entered.
	if answer, qErr := b.Get(bucket, key); err != nil && (qErr != nil || string(answer) != value) {
		return qErr
	}

	return err
}

///TODO: Review this function. Wrote this during session.

//Get is a wrapping function for putting data into the databucket database. It should make queries simple and safe.
func (b *JackalDB) Get(bucket string, query string) (queryReturn []byte, err error) {

	err = b.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		queryReturn = b.Get([]byte(query))

		//Leaving this here for debug rn.
		fmt.Printf("The answer is: %s\n", queryReturn)
		return nil
	})

	//This is an info/warning level debug error. This does NOT indicate a failure of any kind. We will make sure that we're not masking another error when writing this result.
	if err == nil && len(queryReturn) < 1 {
		err = errors.New("database query did not result in an answer")
	}

	return
}

func (b *JackalDB) CreateNestedBucket(root, name string) (bucket *bolt.Bucket, err error) {

	err = b.Db.Update(func(tx *bolt.Tx) error {
		base, err := tx.CreateBucketIfNotExists([]byte(root))

		if err != nil {
			return err
		}

		bucket, err = base.CreateBucketIfNotExists([]byte(name))

		return err
	})

	return
}

///TODO: Implement database functions.
