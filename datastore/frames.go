package datastore

type Frames struct {
	backend DatastoreBackend
}

// NewFrames creates a new Frames data store.
func NewFrames(backendName string, backendConfig any) (*Frames, error) {
	be, err := NewBackend(backendName, backendConfig)
	if err != nil {
		return nil, err
	}

	return &Frames{backend: be}, nil
}

// Get returns a data value for a specific frame from the store .
func (s *Frames) Get(frame int, key string) (string, bool, error) {
	return s.backend.Get(frame, key)
}

// GetKeys returns all the keys for a specific frame from the store.
func (s *Frames) GetKeys(frame int) ([]string, bool, error) {
	return s.backend.GetKeys(frame)
}

// Set sets a key/value for a specific frame in the store.
func (s *Frames) Set(frame int, key, val string) error {
	return s.backend.Set(frame, key, val)
}

// Delete deletes data for a specific frame from the store.
func (s *Frames) Delete(frame int, key string) error {
	return s.backend.Delete(frame, key)
}

// DeleteAll deletes all data for a specific frame from the store.
func (s *Frames) DeleteAll(frame int) error {
	return s.backend.DeleteAll(frame)
}

// Exists returns true if there is any data for a specific frame in the store.
func (s *Frames) Exists(frame int) (bool, error) {
	return s.backend.Exists(frame)

}

// Close releases all storage backend resources.
func (s *Frames) Close() error {
	return s.backend.Close()
}
