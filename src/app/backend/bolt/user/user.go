package user

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/wzt3309/k8sconsole/src/app/backend/bolt/internal"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
	userApi "github.com/wzt3309/k8sconsole/src/app/backend/user/api"
)

const (
	// BucketName represents the name of the bucket where Service store data
	BucketName = "users"
)

// Implements UserService interface.
// Use boltdb to store user information
type service struct {
	db *bolt.DB
}

// NewUserService create a new instance of a service
func NewUserService(db *bolt.DB) (userApi.UserService, error) {
	err := internal.CreateBucket(db, BucketName)
	if err != nil {
		return nil, err
	}

	return &service{
		db: db,
	}, nil
}

// User returns a user by ID
func (self *service) User(ID userApi.UserID) (*userApi.User, error) {
	var user userApi.User
	identifier := internal.Itob(int(ID))

	err := internal.GetObject(self.db, BucketName, identifier, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserByUsername returns a user by username.
func (self *service) UserByUsername(Username string) (*userApi.User, error) {
	var user *userApi.User

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var u userApi.User
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
func (self *service) Users() ([]userApi.User, error) {
	var users = make([]userApi.User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user userApi.User

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
func (self *service) UsersByRole(role userApi.UserRole) ([]userApi.User, error) {
	var users = make([]userApi.User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user userApi.User

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
func (self *service) UpdateUser(ID userApi.UserID, user *userApi.User) error {
	identifier := internal.Itob(int(ID))
	return internal.UpdateObject(self.db, BucketName, identifier, user)
}

// CreateUser creates a new user.
func (self *service) CreateUser(user *userApi.User) error {
	return self.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		user.ID = userApi.UserID(id)

		data, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return bucket.Put(internal.Itob(int(user.ID)), data)
	})
}

// DeleteUser deletes a user.
func (self *service) DeleteUser(ID userApi.UserID) error {
	identifier := internal.Itob(int(ID))
	return internal.DeleteObject(self.db, BucketName, identifier)
}
