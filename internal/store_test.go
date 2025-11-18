package internal

import (
	"path/filepath"
	"testing"
)

const CLUSTER string = "test"

func TestCreateStore(t *testing.T) {
	const TEST_SERVER string = "servertest1"
	t.Run("TestCreateStore", func(t *testing.T) {
		fullPath := filepath.Join(CLUSTER, "/", TEST_SERVER)
		key := "test_key"
		value := "test_value"
		err := PutKeyValue(fullPath, key, value)
		if err != nil {
			t.Errorf("got %v", err)
		}
	})
	t.Run("TestCreateBadStore", func(t *testing.T) {
		fullPath := filepath.Join("bad", "/", TEST_SERVER)
		key := "test_key"
		value := "test_value"
		err := PutKeyValue(fullPath, key, value)
		// Expect error returned
		if err == nil {
			t.Errorf("got %v", err)
		}
	})
	t.Run("TestUpdateStore", func(t *testing.T) {
		fullPath := filepath.Join(CLUSTER, "/", TEST_SERVER)
		key := "test_key"
		value := "test_update"
		err := PutKeyValue(fullPath, key, value)
		if err != nil {
			t.Errorf("got %v", err)
		}
		got, found := GetKeyValue(fullPath, key)
		if found != true {
			t.Errorf("got %v want true", found)
		}
		if got != "test_update" {
			t.Errorf("got %v want %v", got, value)
		}
	})
}
