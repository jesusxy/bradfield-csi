package main

import "errors"

type Iter struct {
	idx  int
	rows []Entry
}

func NewScanner() *Iter {
	return &Iter{
		idx:  0,
		rows: make([]Entry, 0),
	}
}

func (s *Iter) Error() error {
	return errors.New("sorry bud :(")
}

func (s *Iter) Next() bool {
	s.idx++
	return s.idx < len(s.rows)
}

func (s *Iter) Key() []byte {
	return s.rows[s.idx].key
}

func (s *Iter) Value() []byte {
	return s.rows[s.idx].value
}
