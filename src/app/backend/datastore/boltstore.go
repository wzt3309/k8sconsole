package datastore

import (
	"github.com/boltdb/bolt"
	"github.com/wzt3309/k8sconsole/src/app/backend/bolt/user"
	"github.com/wzt3309/k8sconsole/src/app/backend/file"
	"path"
	"time"
)

const (
	databaseFileName = "k8sconsole.db"
)

// Store defines the implementation of datastore.Datastore
// using boltdb as the storage system
type Store struct {
	path        string
	db          *bolt.DB
	FileService file.FileService
	UserService *user.Service
}

// NewStore initializes a new store and the associated services
func NewStore(storePath string, fileService file.FileService) (*Store, error) {
	store := &Store{
		path:        storePath,
		FileService: fileService,
	}

	return store, nil
}

// Opens and initializes the boltdb database.
func (self *Store) Open() error {
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
func (self *Store) Init() error {
	return nil
}

// Closes the boltdb database
func (self *Store) Close() error {
	if self.db != nil {
		return self.db.Close()
	}
	return nil
}

// Initializes the services of store
func (self *Store) initServices() error {
	userService, err := user.NewUserService(self.db)
	if err != nil {
		return err
	}
	self.UserService = userService

	return nil
}
