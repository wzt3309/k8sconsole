package bolt

import (
	"encoding/binary"
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
)

// CreateBucket is a generic function used to create a bucket inside a bolt database.
func CreateBucket(db *bolt.DB, bucketName string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
}

// DeleteObject is a generic function used to delete an object inside a bolt database.
func DeleteObject(db *bolt.DB, bucketName string, key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		return bucket.Delete(key)
	})
}

// GetNextIdentifier is a generic function that returns the specified bucket identifier incremented by 1.
func GetNextIdentifier(db *bolt.DB, bucketName string) int {
	var identifier int

	// find the last identifier of the bucket
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		id := bucket.Sequence()
		identifier = int(id)
		return nil
	})

	identifier++
	return identifier
}

// GetObject is a generic function used to retrieve an unmarshal object from a bolt database
func GetObject(db *bolt.DB, bucketName string, key []byte, object interface{}) error {
	var data []byte

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		val := bucket.Get(key)

		if val == nil {
			return errors.ErrObjectNotFound
		}

		data = make([]byte, len(val))
		copy(data, val)

		return nil
	})

	if err != nil {
		return err
	}

	return json.Unmarshal(data, object)
}

// Itob returns an 8-byte Big-Endian representation of v.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// UpdateObject is a generic function used to update an object inside a bolt database
func UpdateObject(db *bolt.DB, bucketName string, key []byte, object interface{}) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		data, err := json.Marshal(object)
		if err != nil {
			return err
		}

		err = bucket.Put(key, data)
		if err != nil {
			return err
		}
		return nil
	})
}