package level_db

import (
	"encoding/csv"
	"os"
)

type DB interface {
	// Get gets the value for the given key. It returns an error if the
	// DB does not contain the key.
	Get(key []byte) (value []byte, err error)
	// Delete deletes the value for the given key.
	Delete(key []byte) error
	// Has returns true if the DB contains the given key.
	Has(key []byte) (bool, error)
	// Put sets the value for the given key. It overwrites any previous value
	// for that key; a DB is not a multi-map.
	Put(key, value []byte) error
}

type Iterator interface {
	// Error returns any accumulated error. Exhausting all the key/value pairs
	// is not considered to be an error.
	Error() error
	// Next moves the iterator to the next key/value pair.
	// It returns false if the iterator is exhausted.
	Next() bool
	// Key returns the key of the current key/value pair, or nil if done.
	Key() []byte
	// Value returns the value of the current key/value pair, or nil if done.
	Value() []byte
}

func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}
