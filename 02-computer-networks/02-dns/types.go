package main

type Question struct {
	QName  string
	QType  uint16
	QClass uint16
}

type Header struct {
	ID               uint16
	Flags            uint16
	Questions        uint16
	AnswerRecords    uint16
	AuthorityRecords uint16
	AdditionalRecord uint16
}

type ResourceRecord struct {
	Name     uint16
	Type     uint16
	Class    uint16
	TTL      uint16
	RDLength uint16
	RData    uint32
}

type DnsMessage struct {
	Header     Header
	Questions  []Question
	Answers    []ResourceRecord
	Authority  []ResourceRecord
	Additional []ResourceRecord
}
