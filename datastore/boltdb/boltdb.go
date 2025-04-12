package boltdb

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/wasmvision/wasmvision/datastore"
)

type datastoreBoltDB struct {
	db *bolt.DB
}

func newDatastoreBoltDB(args any) (datastore.DatastoreBackend, error) {

	opts, ok := args.(map[string]any)
	if !ok {
		return nil, datastore.ErrInvalidConfigType
	}
	fname, ok := opts["filename"].(string)
	if !ok {
		return nil, fmt.Errorf("%w: filename must be string type", datastore.ErrInvalidConfigType)
	}

	fmode, ok := opts["filemode"].(os.FileMode)
	if !ok {
		return nil, fmt.Errorf("%w: filemode must be os.FileMode type", datastore.ErrInvalidConfigType)
	}

	db, err := bolt.Open(fname, fmode, nil)
	if err != nil {
		return nil, err
	}

	return &datastoreBoltDB{
		db: db,
	}, nil
}

// Get returns a data value from the store.
func (ds *datastoreBoltDB) Get(id any, key string) (string, bool, error) {
	var realID string

	switch v := id.(type) {
	case int:
		realID = strconv.Itoa(v)

	case string:
		realID = v

	default:
		return "", false, datastore.ErrInvalidIDType
	}

	value := ""

	err := ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(realID))
		if b == nil {
			return errors.New("invalid bucket")
		}
		value = string(b.Get([]byte(key)))
		return nil
	})

	return value, value != "", err
}

// GetKeys returns all the keys for a specific id from the store.
func (ds *datastoreBoltDB) GetKeys(id any) ([]string, bool, error) {
	var realID string

	switch v := id.(type) {
	case int:
		realID = strconv.Itoa(v)

	case string:
		realID = v

	default:
		return nil, false, datastore.ErrInvalidIDType
	}

	var keys []string

	err := ds.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(realID))
		if b == nil {
			return errors.New("invalid bucket")
		}
		err := b.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
		return err
	})
	return keys, len(keys) > 0, err
}

// Set sets a config value in the store.
func (ds *datastoreBoltDB) Set(id any, key string, val string) error {
	var realID string

	switch v := id.(type) {
	case int:
		realID = strconv.Itoa(v)

	case string:
		realID = v

	default:
		return datastore.ErrInvalidIDType
	}

	err := ds.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(realID))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(val))
	})

	return err
}

// Delete deletes data from the store.
func (ds *datastoreBoltDB) Delete(id any, key string) error {
	var realID string

	switch v := id.(type) {
	case int:
		realID = strconv.Itoa(v)

	case string:
		realID = v

	default:
		return datastore.ErrInvalidIDType
	}

	err := ds.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(realID))
		if err != nil {
			return err
		}
		return b.Delete([]byte(key))
	})

	return err
}

// DeleteAll deletes all data for a specific id from the store.
func (ds *datastoreBoltDB) DeleteAll(id any) error {
	var realID string

	switch v := id.(type) {
	case int:
		realID = strconv.Itoa(v)

	case string:
		realID = v

	default:
		return datastore.ErrInvalidIDType
	}

	err := ds.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(realID))
	})

	return err
}

// Exists returns true if there is any data for a specific id in the store.
func (ds *datastoreBoltDB) Exists(id any) (bool, error) {
	var realID string

	switch v := id.(type) {
	case int:
		realID = strconv.Itoa(v)

	case string:
		realID = v

	default:
		return false, datastore.ErrInvalidIDType
	}

	exists := false

	err := ds.db.View(func(tx *bolt.Tx) error {
		exists = tx.Bucket([]byte(realID)) != nil
		return nil
	})

	return exists, err

}

func (ds *datastoreBoltDB) Close() error {
	return ds.db.Close()
}

func init() {
	datastore.Register("boltdb", newDatastoreBoltDB)
}
