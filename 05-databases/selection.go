package main

func (s *SelectionOperator) Next() bool {
	for {
		s.idx++
		if s.idx >= len(s.rows) {
			return false // no more rows
		}

		if s.predicate.Evaluate(s.rows[s.idx]) {
			return true // found row that matches predicate
		}
	}
}

func (s *SelectionOperator) Execute() Row {
	return s.rows[s.idx]
}

func NewSelectionOperator(rows []Row, predicate Expression) Node {
	return &SelectionOperator{
		rows:      rows,
		idx:       -1,
		predicate: predicate,
	}
}
