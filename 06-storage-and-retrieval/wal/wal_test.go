package wal

import (
	"os"
	"reflect"
	"testing"
)

func TestWALWriteRead(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "wal")

	if err != nil {
		t.Fatal("failed to create temp WAL file", err)
	}

	tmpFilePath := tmpFile.Name()

	if err := tmpFile.Close(); err != nil {
		t.Fatal("failed to close temp WAL file", err)

	}

	defer func() {
		if err := os.Remove(tmpFilePath); err != nil {
			t.Error("failed to remove file:", err)
		}
	}()

	// create wal
	wal, err := NewWAL(tmpFilePath)

	if err != nil {
		t.Fatalf("failed to create WAL: %v", err)
	}

	// DB operations
	var operations = []Entry{
		{Op: OP_PUT, Key: []byte("The Dark Knight"), KeyLen: uint16(len([]byte("The Dark Knight"))), Val: []byte("Action"), ValLen: uint16(len([]byte("Action")))},
		{Op: OP_PUT, Key: []byte("Casino"), KeyLen: uint16(len([]byte("Casino"))), Val: []byte("Drama"), ValLen: uint16(len([]byte("Drama")))},
		{Op: OP_DELETE, Key: []byte("Casino"), KeyLen: uint16(len([]byte("Casino")))},
	}

	for _, op := range operations {
		if op.Op == OP_PUT {
			err := wal.Put(op.Key, op.Val)
			if err != nil {
				t.Fatalf("Error writing PUT to log: %v", err)
			}
		} else if op.Op == OP_DELETE {
			err := wal.Delete(op.Key)
			if err != nil {
				t.Fatalf("Error DELETE from log: %v", err)
			}
		}
	}

	// reuse WAL
	// read back the entries
	entries, err := wal.ReadEntries()
	if err != nil {
		t.Fatalf("Error reading records from WAL: %v", err)
	}

	// verify that entries match
	if len(entries) != len(operations) {
		t.Fatalf("Expected %d entries, got %d", len(entries), len(operations))
	}

	for i, entry := range entries {
		expectedOp := &operations[i]
		if !reflect.DeepEqual(expectedOp, entry) {
			t.Errorf("Entry %d does not match. Expected %+v, got %+v", i, expectedOp, entry)
		}
	}
}
