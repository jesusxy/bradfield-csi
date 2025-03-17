package main

import (
	"compiler/scanner"
	"testing"
)

func TestParser(t *testing.T) {
	for _, testCase := range []struct {
		input    string
		expected string
	}{
		{
			"alice AND bob",
			"AND(TERM(alice), TERM(bob))",
		},
		{
			"(alice OR bob) and (carol AND (dave or eve))",
			"AND(OR(TERM(alice), TERM(bob)), AND(TERM(carol), OR(TERM(dave), TERM(eve))))",
		},
	} {
		scanner := scanner.NewScanner(string(testCase.input))
		tokens := scanner.ScanTokens()

		parser := NewParser(tokens)
		expression := parser.parseQuery()

		actual := expression.String()

		if testCase.expected != actual {
			t.Errorf("Expected %s, got %s", testCase.expected, actual)
		}
	}
}
