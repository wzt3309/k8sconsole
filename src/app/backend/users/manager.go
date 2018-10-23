package users

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	internal "github.com/wzt3309/k8sconsole/src/app/backend/bolt"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
	userApi "github.com/wzt3309/k8sconsole/src/app/backend/users/api"
)

// Service represents a service to manage user accounts
type Service struct {
	db *bolt.DB
}

// NewService create a new instance of a service
func NewService(db *bolt.DB) (*Service, error) {
	err := internal.CreateBucket(db, userApi.BucketName)
	if err != nil {
		return nil, err
	}

	return &Service{
		db: db,
	}, nil
}

// User returns a user by ID
func (self *Service) User(ID userApi.UserID) (*userApi.User, error) {
	var user userApi.User
	identifier := internal.Itob(int(ID))

	err := internal.GetObject(self.db, userApi.BucketName, identifier, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserByUsername returns a user by username.
func (self *Service) UserByUsername	(Username string) (*userApi.User, error) {
	var user *userApi.User

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userApi.BucketName))
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
func (self *Service) Users() ([]userApi.User, error) {
	var users = make([]userApi.User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userApi.BucketName))
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
func (self *Service) UsersByRole(role userApi.UserRole) ([]userApi.User, error) {
	var users = make([]userApi.User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userApi.BucketName))
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
func (self *Service) UpdateUser(ID userApi.UserID, user *userApi.User) error {
	identifier := internal.Itob(int(ID))
	return internal.UpdateObject(self.db, userApi.BucketName, identifier, user)
}

// CreateUser creates a new user.
func (self *Service) CreateUser(user *userApi.User) error {
	return self.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(userApi.BucketName))

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
func (self *Service) DeleteUser(ID userApi.UserID) error {
	identifier := internal.Itob(int(ID))
	return internal.DeleteObject(self.db, userApi.BucketName, identifier)
}