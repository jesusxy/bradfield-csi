package main

type Row struct {
	Entries []Entry
}

type Entry struct {
	Name  string
	Value string
}

type Node interface {
	Next() bool
	Execute() Row
}

type ScannerOperator struct {
	rows []Row
	idx  int
}

type SelectionOperator struct {
	rows      []Row
	idx       int
	predicate Expression
}
