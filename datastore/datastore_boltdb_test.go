package datastore_test

import (
	"os"
	"testing"

	"github.com/wasmvision/wasmvision/datastore"
	_ "github.com/wasmvision/wasmvision/datastore/boltdb"
	_ "github.com/wasmvision/wasmvision/datastore/frames"
	_ "github.com/wasmvision/wasmvision/datastore/processors"
)

func TestDatastoreDrivers(t *testing.T) {

	cases := []struct {
		driver string
		config any
		bucket any
		key    string
		value  string
	}{
		{driver: "frames", bucket: 1, key: "frames_test_key", value: "frames_test_value", config: nil},
		{driver: "processors", bucket: "test_processor", key: "processor_test_key", value: "processor_test_value", config: nil},
		{driver: "boltdb", bucket: "test_boltdb", key: "boltdb_test_key", value: "boltdb_test_value", config: map[string]any{"filename": "test.db", "filemode": os.FileMode(0644)}},
	}

	for _, c := range cases {
		ds, err := datastore.New(c.driver, c.config)
		if err != nil {
			t.Error(err)
		}

		err = ds.Set(c.bucket, c.key, c.value)
		if err != nil {
			t.Error(err)
		}

		gotValue, found, err := ds.Get(c.bucket, c.key)
		if err != nil {
			t.Error(err)
		}
		if !found {
			t.Error("value not found")
		}

		if gotValue != c.value {
			t.Errorf("got %s want %s", gotValue, c.value)
		}

		keys, _, err := ds.GetKeys(c.bucket)
		if err != nil {
			t.Error(err)
		}
		if keys[0] != c.key {
			t.Errorf("got %s want %s", keys[0], c.key)
		}

		ds.Close()
	}
}
