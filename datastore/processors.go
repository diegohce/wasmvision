package datastore

type Processors struct {
	backend DatastoreBackend
}

// NewProcessors creates a new Processors data store.
func NewProcessors(backendName string, backendConfig any) (*Processors, error) {

	be, err := NewBackend(backendName, backendConfig)
	if err != nil {
		return nil, err
	}

	return &Processors{backend: be}, nil
}

// Get returns a data value from the store.
func (s *Processors) Get(processor string, key string) (string, bool, error) {
	return s.backend.Get(processor, key)
}

// GetKeys returns all the keys for a specific processor from the store.
func (s *Processors) GetKeys(processor string) ([]string, bool, error) {
	return s.backend.GetKeys(processor)
}

// Set sets a config value in the store.
func (s *Processors) Set(processor string, key string, val string) error {
	return s.backend.Set(processor, key, val)
}

// Delete deletes data from the store.
func (s *Processors) Delete(processor string, key string) error {
	return s.backend.Delete(processor, key)
}

// DeleteAll deletes all data for a specific processor from the store.
func (s *Processors) DeleteAll(processor string) error {
	return s.backend.DeleteAll(processor)
}

// Exists returns true if there is any data for a specific processor in the store.
func (s *Processors) Exists(processor string) (bool, error) {
	return s.backend.Exists(processor)
}

// Close releases all storage backend resources.
func (s *Processors) Close() error {
	return s.backend.Close()
}
