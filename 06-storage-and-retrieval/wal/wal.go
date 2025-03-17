package wal

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"os"
)

const (
	OP_PUT    = 0x01
	OP_DELETE = 0x02
)

type WAL struct {
	file *os.File
	path string
}

func NewWAL(filePath string) (*WAL, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	return &WAL{
		file: file,
		path: filePath,
	}, nil
}

// this will be called by top level db
func (wal *WAL) Put(key, val []byte) error {
	entry := Entry{
		Op:     OP_PUT,
		KeyLen: uint16(len(key)),
		Key:    key,
		ValLen: uint16(len(val)),
		Val:    val,
	}

	return wal.writeEntry(entry)
}

// this will be called by top level db
func (wal *WAL) Delete(key []byte) error {
	entry := Entry{
		Op:     OP_DELETE,
		KeyLen: uint16(len(key)),
		Key:    key,
	}

	return wal.writeEntry(entry)
}

func (wal *WAL) writeEntry(entry Entry) error {
	// serialize data
	data := entry.serialize()

	// write serialized data to the WAL
	_, err := wal.file.Write(data)

	return err
}

// for each byte slice from the file convert it to an Entry struct and append to entries []
func (wal *WAL) ReadEntries() ([]*Entry, error) {
	var entries []*Entry

	// open the file where WAL is
	fil, err := os.Open(wal.path)

	if err != nil {
		return nil, err
	}

	// defer walClose()
	defer fil.Close()

	// create buffered reader
	reader := bufio.NewReader(fil)

	for {
		// deserialize data
		entry, err := deserialize(reader)

		if err != nil {
			if err == io.EOF { // if end of file, break out of loop
				break
			}
			return nil, err
		}
		// append to entries
		entries = append(entries, entry)
	}

	return entries, nil
}

func (wal *WAL) Close() error {
	return wal.file.Close()
}

// WAL Entry
type Entry struct {
	Op     byte
	KeyLen uint16
	Key    []byte
	ValLen uint16
	Val    []byte
}

// TODO: add error handling
func (e *Entry) serialize() []byte {
	buffer := bytes.Buffer{}
	// serialze op
	buffer.WriteByte(e.Op)
	// serialize key, kelen
	binary.Write(&buffer, binary.LittleEndian, e.KeyLen)
	buffer.Write(e.Key)

	// serialize val, vallen ONLY for PUT
	if e.Op == OP_PUT {
		binary.Write(&buffer, binary.LittleEndian, e.ValLen)
		buffer.Write(e.Val)
	}

	return buffer.Bytes()
}

func deserialize(reader *bufio.Reader) (*Entry, error) {
	entry := &Entry{}

	op, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	entry.Op = op

	keyLenBytes := make([]byte, 2)
	if _, err := io.ReadFull(reader, keyLenBytes); err != nil {
		return nil, err
	}

	keyLen := binary.LittleEndian.Uint16(keyLenBytes)
	entry.KeyLen = keyLen

	key := make([]byte, keyLen)
	if _, err := io.ReadFull(reader, key); err != nil {
		return nil, err
	}

	entry.Key = key

	// for PUT operations, deserialize val
	if entry.Op == OP_PUT {
		valLenBytes := make([]byte, 2)
		if _, err := io.ReadFull(reader, valLenBytes); err != nil {
			return nil, err
		}

		valLen := binary.LittleEndian.Uint16(valLenBytes)
		entry.ValLen = valLen

		val := make([]byte, valLen)
		if _, err := io.ReadFull(reader, val); err != nil {
			return nil, err
		}

		entry.Val = val

	}

	return entry, nil
}
