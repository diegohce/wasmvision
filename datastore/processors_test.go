package datastore_test

import (
	"slices"
	"testing"

	"github.com/wasmvision/wasmvision/datastore"
	_ "github.com/wasmvision/wasmvision/datastore/memory"
)

func TestProcessors(t *testing.T) {
	t.Run("get", func(t *testing.T) {

		s, _ := datastore.NewProcessors("memory", nil)

		s.Set("proc", "key", "value")

		val, ok, _ := s.Get("proc", "key")
		if !ok {
			t.Errorf("key not found")
		}

		if string(val) != "value" {
			t.Errorf("unexpected value: %s", val)
		}
	})

	t.Run("exists", func(t *testing.T) {
		s, _ := datastore.NewProcessors("memory", nil)

		s.Set("proc", "key", "value")

		ok, _ := s.Exists("proc")
		if !ok {
			t.Errorf("not found")
		}
	})

	t.Run("getKeys", func(t *testing.T) {

		s, _ := datastore.NewProcessors("memory", nil)

		s.Set("proc", "key", "value")
		s.Set("proc", "key2", "value2")
		s.Set("proc", "key3", "value3")

		keys, ok, _ := s.GetKeys("proc")
		if !ok {
			t.Errorf("processor not found")
		}

		if len(keys) != 3 {
			t.Errorf("unexpected number of keys: %d", len(keys))
		}

		if !slices.Contains(keys, "key") {
			t.Errorf("key not found")
		}
	})

	t.Run("set", func(t *testing.T) {
		s, _ := datastore.NewProcessors("memory", nil)

		err := s.Set("proc", "key", "value")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		val, ok, _ := s.Get("proc", "key")
		if !ok {
			t.Errorf("key not found")
		}

		if string(val) != "value" {
			t.Errorf("unexpected value: %s", val)
		}
	})

	t.Run("delete", func(t *testing.T) {
		s, _ := datastore.NewProcessors("memory", nil)

		s.Set("proc", "key", "value")

		s.Delete("proc", "key")

		_, ok, _ := s.Get("proc", "key")
		if ok {
			t.Errorf("key not deleted")
		}
	})
}
