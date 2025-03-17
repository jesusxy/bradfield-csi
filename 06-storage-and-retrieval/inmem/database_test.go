package main

import (
	"level_db"
	"reflect"
	"testing"
)

func seedDbFromCSV(db level_db.DB, filePath string) error {
	records, err := level_db.ReadCSV(filePath)

	if err != nil {
		return err
	}

	for _, record := range records {
		key := []byte(record[0])
		value := []byte(record[1])

		if err := db.Put(key, value); err != nil {
			return err
		}
	}

	return nil

}
func TestDbOperations(t *testing.T) {
	memdb := NewMemDb()
	err := seedDbFromCSV(memdb, "./movies_new.csv")

	if err != nil {
		t.Fatalf("Failed to seed database from CSV: %v", err)
	}

	// Test Has
	t.Run("Has", func(t *testing.T) {
		key := []byte("27")
		exists, err := memdb.Has(key)

		if err != nil || !exists {
			t.Errorf("Has() error = %v, exists = %v, want = %v", err, exists, true)
		}
	})

	// Test Get
	t.Run("Get", func(t *testing.T) {
		key, value := []byte("16"), []byte("Casino (1995)")
		got, err := memdb.Get(key)

		if err != nil || !reflect.DeepEqual(got, value) {
			t.Errorf("Get() got = %v, want %v", string(got), string(value))
		}
	})

	// Test Put
	t.Run("Put", func(t *testing.T) {
		key, value := []byte("201"), []byte("Entourage")
		if err := memdb.Put(key, value); err != nil {
			t.Errorf("Put() error = %v, want = %v", err, false)
		}
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		key := []byte("201")
		if err := memdb.Delete(key); err != nil {
			t.Errorf("Delete() err = %v, want = %v", err, false)
		}
	})

	// Test Verify Delete
	t.Run("VerifyDeletion", func(t *testing.T) {
		key := []byte("201")
		exists, _ := memdb.Has(key)

		if exists {
			t.Errorf("Has() after Delete() exists = %v, want = %v", exists, false)
		}
	})
}
