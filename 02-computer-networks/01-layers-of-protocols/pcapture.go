package main

import (
	"fmt"
	"io"
	"os"
)

const (
	GlobalFileHeaderLength = 24
	PacketHeaderLength     = 16
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("net.cap")
	check(err)

	defer file.Close()

	// parse global header
	ghbuf := make([]byte, GlobalFileHeaderLength)
	_, err = file.Read(ghbuf)
	check(err)

	gfh, err := ParseGlobalFileHeader(ghbuf)

	totalPackets := 0

	// loop through packets

	for {
		phbuf := make([]byte, PacketHeaderLength)
		_, err = file.Read(phbuf)

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		ph, err := ParsePacketHeader(phbuf)
		check(err)

		pktdata := make([]byte, ph.CapturedLen)
		_, err = file.Read(pktdata)
		check(err)

		totalPackets++
	}

	gfh.Print()
	fmt.Println("total packets read: ", totalPackets)
}
