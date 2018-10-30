package user

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/bolt/internal"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
)

const (
	// BucketName represents the name of the bucket where service store data
	BucketName = "users"
)

// Implements UserService interface.
// Use boltdb to store user information
type service struct {
	db *bolt.DB
}

// NewUserService create a new instance of a service
func NewUserService(db *bolt.DB) (api.UserService, error) {
	err := internal.CreateBucket(db, BucketName)
	if err != nil {
		return nil, err
	}

	return &service{
		db: db,
	}, nil
}

// User returns a user by ID
func (self *service) User(ID api.UserID) (*api.User, error) {
	var user api.User
	identifier := internal.Itob(int(ID))

	err := internal.GetObject(self.db, BucketName, identifier, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserByUsername returns a user by username.
func (self *service) UserByUsername(Username string) (*api.User, error) {
	var user *api.User

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var u api.User
			err := json.Unmarshal(v, &u)
			if err != nil {
				return err
			}

			if u.Username == Username {
				user = &u
				break
			}
		}

		if user == nil {
			return errors.ErrObjectNotFound
		}
		return nil
	})

	return user, err
}

// Users return a slice containing all the users.
func (self *service) Users() ([]api.User, error) {
	var users = make([]api.User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user api.User

			err := json.Unmarshal(v, &user)
			if err != nil {
				return err
			}

			users = append(users, user)
		}
		return nil
	})

	return users, err
}

// UsersByRole return an slice contains all the users with the specified role.
func (self *service) UsersByRole(role api.UserRole) ([]api.User, error) {
	var users = make([]api.User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user api.User

			err := json.Unmarshal(v, &user)
			if err != nil {
				return err
			}

			if user.Role == role {
				users = append(users, user)
			}
		}

		return nil
	})

	return users, err
}

// UpdateUser update old user
func (self *service) UpdateUser(ID api.UserID, user *api.User) error {
	identifier := internal.Itob(int(ID))
	return internal.UpdateObject(self.db, BucketName, identifier, user)
}

// CreateUser creates a new user.
func (self *service) CreateUser(user *api.User) error {
	return self.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		user.ID = api.UserID(id)

		data, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return bucket.Put(internal.Itob(int(user.ID)), data)
	})
}

// DeleteUser deletes a user.
func (self *service) DeleteUser(ID api.UserID) error {
	identifier := internal.Itob(int(ID))
	return internal.DeleteObject(self.db, BucketName, identifier)
}
