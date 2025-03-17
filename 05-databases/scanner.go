package main

func (s *ScannerOperator) Next() bool {
	s.idx++
	return s.idx < len(s.rows)
}

func (s *ScannerOperator) Execute() Row {
	return s.rows[s.idx]
}

func NewScannerOperator(rows []Row) Node {
	return &ScannerOperator{
		rows: rows,
		idx:  -1,
	}
}
