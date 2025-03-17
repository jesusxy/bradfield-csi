package dbclone

import (
	"bytes"
	"level_db"
	"log"
	"os"
	"testing"
)

func TestDatabasePutAndGet(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "wal")
	if err != nil {
		t.Fatal("failed to create temp WAL file", err)
	}

	tmpFilePath := tmpFile.Name()

	if err := tmpFile.Close(); err != nil {
		t.Fatal("failed to close tmp wal file", err)
	}

	defer func() {
		if err := os.Remove(tmpFilePath); err != nil {
			t.Error("failed to remove file", err)
		}
	}()

	// create db
	db, err := NewCloneDB(tmpFilePath)
	if err != nil {
		t.Fatal("failed to db instance with temp WAL", err)
	}

	movies, err := level_db.ReadCSV("../movies_fifty.csv")
	if err != nil {
		t.Fatalf("failed to read dataset from csv")
	}

	t.Run("PUT operations", func(t *testing.T) {
		for _, movie := range movies[:26] {
			key := []byte(movie[1])
			val := []byte(movie[2])
			log.Printf("Inserting: Key=%s, Value=%s", key, val)
			if err := db.Put(key, val); err != nil {
				t.Errorf("failed to put key value pair: %v", err)
			}
		}
	})

	// Test GET operation
	t.Run("GET operation", func(t *testing.T) {
		targetKey := []byte("Casino (1995)")
		expected := []byte("Crime|Drama")

		got, err := db.Get(targetKey)
		if err != nil {
			t.Errorf("Failed to get value: %v", err)
		}

		if !bytes.Equal(got, expected) {
			t.Errorf("GET got %s, expected %s", got, expected)
		}
	})
}
