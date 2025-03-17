package main

import (
	"errors"
)

type MemDb struct {
	records map[string][]byte
}

type Entry struct {
	key   []byte
	value []byte
}

func NewMemDb() *MemDb {
	return &MemDb{
		records: make(map[string][]byte),
	}
}

func (memdb *MemDb) Delete(key []byte) error {
	delete(memdb.records, string(key))
	return nil
}

func (memdb *MemDb) Get(key []byte) (value []byte, err error) {
	entry, ok := memdb.records[string(key)]

	if !ok {
		return nil, errors.New("no record found for given key")
	}

	return entry, nil
}

func (memdb *MemDb) Has(key []byte) (bool, error) {
	_, ok := memdb.records[string(key)]

	if !ok {
		return false, nil
	}

	return true, nil
}

func (memdb *MemDb) Put(key, value []byte) error {
	memdb.records[string(key)] = value
	return nil
}
