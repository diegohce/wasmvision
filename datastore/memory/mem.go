package memory

import "github.com/wasmvision/wasmvision/datastore"

type datastoreMemory struct {
	m map[any]map[string]string
}

func newDatastoreMemory(_ any) (datastore.DatastoreBackend, error) {
	return &datastoreMemory{
		m: make(map[any]map[string]string),
	}, nil
}

// Get returns a data value from the store.
func (ds *datastoreMemory) Get(id any, key string) (string, bool, error) {

	col, ok := ds.m[id]
	if !ok {
		return "", false, nil
	}

	val, ok := col[key]
	return val, ok, nil
}

// GetKeys returns all the keys for a specific id from the store.
func (ds *datastoreMemory) GetKeys(id any) ([]string, bool, error) {
	col, ok := ds.m[id]
	if !ok {
		return nil, false, nil
	}

	keys := make([]string, 0, len(col))
	for k := range col {
		keys = append(keys, k)
	}

	return keys, true, nil

}

// Set sets a config value in the store.
func (ds *datastoreMemory) Set(id any, key string, val string) error {
	col, ok := ds.m[id]
	if !ok {
		ds.m[id] = make(map[string]string)
		col = ds.m[id]
	}

	col[key] = val
	return nil
}

// Delete deletes data from the store.
func (ds *datastoreMemory) Delete(id any, key string) error {
	col, ok := ds.m[id]
	if ok {
		delete(col, key)
	}
	return nil
}

// DeleteAll deletes all data for a specific id from the store.
func (ds *datastoreMemory) DeleteAll(id any) error {
	col, ok := ds.m[id]
	if ok {
		delete(ds.m, id)
		for key := range col {
			delete(col, key)
		}
	}
	return nil
}

// Exists returns true if there is any data for a specific id in the store.
func (ds *datastoreMemory) Exists(id any) (bool, error) {
	_, ok := ds.m[id]
	return ok, nil
}

func (ds *datastoreMemory) Close() error { return nil }

func init() {
	datastore.Register("memory", newDatastoreMemory)
}
