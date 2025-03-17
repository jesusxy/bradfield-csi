package main

import "strings"

type Expression interface {
	Evaluate(Row) bool
}

type EqualsExpression struct {
	Column string
	Value  string
}

func (exp EqualsExpression) Evaluate(r Row) bool {
	for _, entry := range r.Entries {
		if entry.Name == exp.Column && entry.Value == exp.Value {
			return true
		}
	}

	return false
}

type ContainsGenreExpression struct {
	Column string
	Substr string
}

func (exp ContainsGenreExpression) Evaluate(r Row) bool {
	for _, entry := range r.Entries {
		if entry.Name == exp.Column && strings.Contains(entry.Value, exp.Substr) {
			return true
		}
	}

	return false
}
