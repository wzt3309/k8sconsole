package user

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/wzt3309/k8sconsole/src/app/backend/bolt/internal"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
)

// Service represents a service to manage user accounts
type Service struct {
	db *bolt.DB
}

// NewService create a new instance of a service
func NewService(db *bolt.DB) (*Service, error) {
	err := internal.CreateBucket(db, BucketName)
	if err != nil {
		return nil, err
	}

	return &Service{
		db: db,
	}, nil
}

// User returns a user by ID
func (self *Service) User(ID UserID) (*User, error) {
	var user User
	identifier := internal.Itob(int(ID))

	err := internal.GetObject(self.db, BucketName, identifier, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserByUsername returns a user by username.
func (self *Service) UserByUsername(Username string) (*User, error) {
	var user *User

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var u User
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
func (self *Service) Users() ([]User, error) {
	var users = make([]User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user User

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
func (self *Service) UsersByRole(role UserRole) ([]User, error) {
	var users = make([]User, 0)

	err := self.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user User

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
func (self *Service) UpdateUser(ID UserID, user *User) error {
	identifier := internal.Itob(int(ID))
	return internal.UpdateObject(self.db, BucketName, identifier, user)
}

// CreateUser creates a new user.
func (self *Service) CreateUser(user *User) error {
	return self.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		user.ID = UserID(id)

		data, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return bucket.Put(internal.Itob(int(user.ID)), data)
	})
}

// DeleteUser deletes a user.
func (self *Service) DeleteUser(ID UserID) error {
	identifier := internal.Itob(int(ID))
	return internal.DeleteObject(self.db, BucketName, identifier)
}
