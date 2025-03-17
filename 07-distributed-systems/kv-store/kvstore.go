package main

import (
	"encoding/json"
	"os"
)

type KeyValStore struct {
	FilePath string
	Store    map[string]string
}

func NewKeyValStore(filepath string) (*KeyValStore, error) {
	kvs := &KeyValStore{
		FilePath: filepath,
		Store:    make(map[string]string),
	}

	// Attempt to load data from disk, ignore error if the file doesn't exist
	if err := kvs.LoadFromDisk(); err != nil {
		if !os.IsNotExist(err) { // file exists but could not load data
			return nil, err
		}

		// file does not exist, proceed with empty store
	}

	return kvs, nil
}

func (kvs *KeyValStore) SaveToDisk() error {
	file, err := os.Create(kvs.FilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	// encode data
	encoder := json.NewEncoder(file)
	err = encoder.Encode(kvs.Store)

	return err
}

func (kvs *KeyValStore) LoadFromDisk() error {
	// open file
	file, err := os.Open(kvs.FilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	// decode data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&kvs.Store)

	return err
}
