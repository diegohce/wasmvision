package datastore

import "errors"

// DatastoreBackend
//
// id is of type 'any' for its value to be independent from the
// actual backend's implementation of id.
//
// i.e.:
//
// processors: the id is of type string.
//
// frames: the id is of type int.
type DatastoreBackend interface {
	// Get returns a data value from the store.
	Get(id any, key string) (string, bool, error)

	// GetKeys returns all the keys for a specific id from the store.
	GetKeys(id any) ([]string, bool, error)

	// Set sets a config value in the store.
	Set(id any, key string, val string) error

	// Delete deletes data from the store.
	Delete(id any, key string) error

	// DeleteAll deletes all data for a specific id from the store.
	DeleteAll(id any) error

	// Exists returns true if there is any data for a specific id in the store.
	Exists(id any) (bool, error)

	// Close releases Datastore underliying resources.
	Close() error
}

type newDatastoreBackendFunc func(config any) (DatastoreBackend, error)

var (
	datastores           = map[string]newDatastoreBackendFunc{}
	ErrInvalidDatastore  = errors.New("invalid datastore")
	ErrInvalidConfigType = errors.New("invalid config type")
	ErrInvalidIDType     = errors.New("invalid id type")
)

func Register(name string, fn newDatastoreBackendFunc) {
	datastores[name] = fn
}

func NewBackend(name string, config any) (DatastoreBackend, error) {
	newDSFn, exists := datastores[name]
	if !exists {
		return nil, ErrInvalidDatastore
	}
	return newDSFn(config)
}
