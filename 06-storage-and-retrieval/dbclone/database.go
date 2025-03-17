package dbclone

import (
	"fmt"
	"level_db"
	"level_db/skiplist"
	"level_db/wal"
	"log"
)

type Database struct {
	memdb  *skiplist.MemDb
	wal    *wal.WAL
	waldir string
}

func NewCloneDB(waldir string) (level_db.DB, error) {
	wal, err := wal.NewWAL(waldir)
	if err != nil {
		log.Printf("Error creating WAL with filepath: %v", waldir)
		return nil, err
	}

	db := &Database{
		memdb:  skiplist.NewMemDb(),
		wal:    wal,
		waldir: waldir,
	}

	// call db.replayWal
	if err := db.replayWal(); err != nil {
		log.Printf("Error during WAL replay: %v", err)
		return nil, err
	}

	return db, nil
}

func (db *Database) Delete(key []byte) error {
	// add k to Wal
	if err := db.wal.Delete(key); err != nil {
		log.Printf("Failed to write DELETE operation to WAL for key %s: %v", key, err)
		return err
	}
	// call memdb.delete
	if err := db.memdb.Delete(key); err != nil {
		log.Printf("Failed to apply DELETE operation to Memdb for key: %s, %v", key, err)
	}

	return nil
}

func (db *Database) Put(key, value []byte) error {
	if err := db.wal.Put(key, value); err != nil {
		log.Printf("Failed to write PUT operation to WAL for key: %s: %v", key, err)
		return err
	}

	if err := db.memdb.Put(key, value); err != nil {
		log.Printf("Failed to apply PUT operation to Memdb for key: %s: %v", key, err)
		return err
	}

	return nil
}

func (db *Database) Has(key []byte) (bool, error) {
	return db.memdb.Has(key)
}

func (db *Database) Get(key []byte) ([]byte, error) {
	record, err := db.memdb.Get(key)
	if err != nil {
		log.Printf("Error retrieving record for given key: %v", err)
		return nil, err
	}

	return record, nil
}

func (db *Database) replayWal() error {
	entries, err := db.wal.ReadEntries()
	if err != nil {
		log.Printf("Error reading records from WAL: %v", err)
		return fmt.Errorf("error reading records from WAL: %v", err)
	}

	if len(entries) == 0 {
		log.Println("No entries in WAL to replay: likely a fresh database")
		return nil
	}

	for _, entry := range entries {
		switch entry.Op {
		case wal.OP_PUT:
			if err := db.memdb.Put(entry.Key, entry.Val); err != nil {
				log.Printf("Error applying PUT during WAL replay: %v", err)
				return err
			}
		case wal.OP_DELETE:
			if err := db.memdb.Delete(entry.Key); err != nil {
				log.Printf("Error applying DELETE during WAL replay: %v", err)
				return err
			}
		default:
			log.Printf("Unsupported operation %d encountered during WAL replay", entry.Op)
		}
	}

	return nil
}
