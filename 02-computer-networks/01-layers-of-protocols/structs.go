package main

import "fmt"

type GlobalFileHeader struct {
	MagicNumber    uint32
	MajorVersion   uint16
	MinorVersion   uint16
	TzOffset       uint32
	TzAccuracy     uint32
	Snapshotlength uint32
	LinkType       uint32
}

func (gfh GlobalFileHeader) Print() {
	fmt.Printf("FileHeader:\nMagicNum: %x\nMajorVer: %d\nMinorVer: %d\nSanpshotLen: %d\nLinkType: %d\n", gfh.MagicNumber, gfh.MajorVersion, gfh.MinorVersion, gfh.Snapshotlength, gfh.LinkType)
}

type PacketHeader struct {
	TimestampSeconds      uint32
	TimestampMicroSeconds uint32
	CapturedLen           uint32
	OriginalLen           uint32
}

func (ph PacketHeader) Print() {
	fmt.Printf("PacketHeader:\nCapturedLength: %d\nOriginalLength:%d", ph.CapturedLen, ph.OriginalLen)
}
