package processors

import "github.com/wasmvision/wasmvision/datastore"

type datastoreProcessors struct {
	p *datastore.Processors
}

func newDatastoreProcessors(_ any) (datastore.Datastore, error) {

	m := make(map[string]map[string]string)
	p := datastore.NewProcessors(m)

	return &datastoreProcessors{
		p: p,
	}, nil
}

// Get returns a data value from the store.
func (ds *datastoreProcessors) Get(id any, key string) (string, bool, error) {
	realID, ok := id.(string)
	if !ok {
		return "", false, datastore.ErrInvalidIDType
	}
	s, b := ds.p.Get(realID, key)
	return s, b, nil
}

// GetKeys returns all the keys for a specific id from the store.
func (ds *datastoreProcessors) GetKeys(id any) ([]string, bool, error) {
	realID, ok := id.(string)
	if !ok {
		return nil, false, datastore.ErrInvalidIDType
	}
	s, b := ds.p.GetKeys(realID)
	return s, b, nil
}

// Set sets a config value in the store.
func (ds *datastoreProcessors) Set(id any, key string, val string) error {
	realID, ok := id.(string)
	if !ok {
		return datastore.ErrInvalidIDType
	}
	ds.p.Set(realID, key, val)
	return nil
}

// Delete deletes data from the store.
func (ds *datastoreProcessors) Delete(id any, key string) error {
	realID, ok := id.(string)
	if !ok {
		return datastore.ErrInvalidIDType
	}
	ds.p.Delete(realID, key)
	return nil
}

// DeleteAll deletes all data for a specific id from the store.
func (ds *datastoreProcessors) DeleteAll(id any) error {
	realID, ok := id.(string)
	if !ok {
		return datastore.ErrInvalidIDType
	}
	ds.p.DeleteAll(realID)
	return nil
}

// Exists returns true if there is any data for a specific id in the store.
func (ds *datastoreProcessors) Exists(id any) (bool, error) {
	realID, ok := id.(string)
	if !ok {
		return false, datastore.ErrInvalidIDType
	}

	return ds.p.Exists(realID), nil
}

func (ds *datastoreProcessors) Close() error { return nil }

func init() {
	datastore.Register("processors", newDatastoreProcessors)
}
