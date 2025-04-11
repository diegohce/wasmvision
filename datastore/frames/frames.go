package frames

import "github.com/wasmvision/wasmvision/datastore"

type datastoreFrames struct {
	f *datastore.Frames
}

func newDatastoreFrames(_ any) (datastore.Datastore, error) {

	m := make(map[int]map[string]string)
	f := datastore.NewFrames(m)

	return &datastoreFrames{
		f: f,
	}, nil
}

// Get returns a data value from the store.
func (ds *datastoreFrames) Get(id any, key string) (string, bool, error) {
	realID, ok := id.(int)
	if !ok {
		return "", false, datastore.ErrInvalidIDType
	}
	s, b := ds.f.Get(realID, key)
	return s, b, nil
}

// GetKeys returns all the keys for a specific id from the store.
func (ds *datastoreFrames) GetKeys(id any) ([]string, bool, error) {
	realID, ok := id.(int)
	if !ok {
		return nil, false, datastore.ErrInvalidIDType
	}
	s, b := ds.f.GetKeys(realID)
	return s, b, nil
}

// Set sets a config value in the store.
func (ds *datastoreFrames) Set(id any, key string, val string) error {
	realID, ok := id.(int)
	if !ok {
		return datastore.ErrInvalidIDType
	}
	ds.f.Set(realID, key, val)
	return nil
}

// Delete deletes data from the store.
func (ds *datastoreFrames) Delete(id any, key string) error {
	realID, ok := id.(int)
	if !ok {
		return datastore.ErrInvalidIDType
	}
	ds.f.Delete(realID, key)
	return nil
}

// DeleteAll deletes all data for a specific id from the store.
func (ds *datastoreFrames) DeleteAll(id any) error {
	realID, ok := id.(int)
	if !ok {
		return datastore.ErrInvalidIDType
	}
	ds.f.DeleteAll(realID)
	return nil
}

// Exists returns true if there is any data for a specific id in the store.
func (ds *datastoreFrames) Exists(id any) (bool, error) {
	realID, ok := id.(int)
	if !ok {
		return false, datastore.ErrInvalidIDType
	}

	return ds.f.Exists(realID), nil
}

func (ds *datastoreFrames) Close() error { return nil }

func init() {
	datastore.Register("frames", newDatastoreFrames)
}
