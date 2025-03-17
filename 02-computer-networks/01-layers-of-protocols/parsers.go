package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

var InvalidHeaderLengthErr error = errors.New("invalid header length")

func ParsePacketHeader(b []byte) (PacketHeader, error) {
	if len(b) != PacketHeaderLength {
		return PacketHeader{}, InvalidHeaderLengthErr
	}

	packetHeader := PacketHeader{
		TimestampSeconds:      binary.LittleEndian.Uint32(b[:4]),
		TimestampMicroSeconds: binary.LittleEndian.Uint32(b[4:8]),
		CapturedLen:           binary.LittleEndian.Uint32(b[8:12]),
		OriginalLen:           binary.LittleEndian.Uint32(b[12:16]),
	}

	if packetHeader.CapturedLen != packetHeader.OriginalLen {
		fmt.Printf("Captured packet length: %d should equal Original packet length: %d", packetHeader.CapturedLen, packetHeader.OriginalLen)
	}

	return packetHeader, nil
}

func ParseGlobalFileHeader(b []byte) (GlobalFileHeader, error) {
	if len(b) != GlobalFileHeaderLength {
		return GlobalFileHeader{}, InvalidHeaderLengthErr
	}

	return GlobalFileHeader{
		MagicNumber:    binary.LittleEndian.Uint32(b[:4]),
		MajorVersion:   binary.LittleEndian.Uint16(b[4:6]),
		MinorVersion:   binary.LittleEndian.Uint16(b[6:8]),
		TzOffset:       binary.LittleEndian.Uint32(b[8:12]),
		TzAccuracy:     binary.LittleEndian.Uint32(b[12:16]),
		Snapshotlength: binary.LittleEndian.Uint32(b[16:20]),
		LinkType:       binary.LittleEndian.Uint32(b[20:24]),
	}, nil
}
