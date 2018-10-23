package bolt

import "github.com/boltdb/bolt"

const (
	dbFileName = "k8sconsole.db"
)

type Store struct {
	path                  string
	db                    *bolt.DB
	checkForDataMigration bool
}
