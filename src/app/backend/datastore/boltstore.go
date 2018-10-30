package datastore

import (
	"github.com/boltdb/bolt"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	boltUser "github.com/wzt3309/k8sconsole/src/app/backend/bolt/user"
	"path"
	"time"
)

const (
	databaseFileName = "k8sconsole.db"
)

// BoltDBStore defines the implementation of datastore.Datastore
// using boltdb as the storage system
type BoltDBStore struct {
	path        string
	db          *bolt.DB
	FileService api.FileService
	UserService api.UserService
}

// NewBoltDBStore initializes a new store and the associated services
func NewBoltDBStore(storePath string, fileService api.FileService) (*BoltDBStore, error) {
	store := &BoltDBStore{
		path:        storePath,
		FileService: fileService,
	}

	return store, nil
}

// Opens and initializes the boltdb database.
func (self *BoltDBStore) Open() error {
	databasePath := path.Join(self.path, databaseFileName)

	db, err := bolt.Open(databasePath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	self.db = db

	return self.initServices()
}

// Initializes the store
// TODO(wzt3309) This method does nothing now
func (self *BoltDBStore) Init() error {
	return nil
}

// Closes the boltdb database
func (self *BoltDBStore) Close() error {
	if self.db != nil {
		return self.db.Close()
	}
	return nil
}

func (self *BoltDBStore) GetUserService() api.UserService {
	return self.UserService
}

// Initializes the services of store
func (self *BoltDBStore) initServices() error {

	// init UserService
	userService, err := boltUser.NewUserService(self.db)
	if err != nil {
		return err
	}
	self.UserService = userService

	return nil
}
