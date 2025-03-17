package scanner

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	LEFT_PAREN = iota
	RIGHT_PAREN
	AND
	OR
	NOT
	IDENTIFIER
	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
}

func (t Token) String() string {
	// TODO: create mapper to get token Type name instead of int
	return fmt.Sprintf("{Type: %d, Lexeme: %s, Literal: %v}", t.Type, t.Lexeme, t.Literal)
}

func NewScanner(input string) *Scanner {
	return &Scanner{
		source: input,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{Type: EOF, Lexeme: "", Literal: nil})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	char := s.advance()

	switch char {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	default:
		if unicode.IsLetter(rune(char)) {
			s.identifier()
		} else if unicode.IsSpace(rune(char)) {
			// skip spaces
		} else {
			fmt.Printf("Unexpected character:%c\n", char)
		}
	}
}

func (s *Scanner) advance() byte {
	s.current++
	// returns next char in input str
	return s.source[s.current-1]
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: text, Literal: literal})
}

func (s *Scanner) identifier() {
	for !s.isAtEnd() && (unicode.IsLetter(rune(s.source[s.current])) || s.source[s.current] == '_') {
		s.current++
	}

	// check if identifier is a reserved word
	text := strings.ToUpper(s.source[s.start:s.current])
	switch text {
	case "AND":
		s.addToken(AND)
	case "OR":
		s.addToken(OR)
	case "NOT":
		s.addToken(NOT)
	default:
		s.addTokenWithLiteral(IDENTIFIER, s.source[s.start:s.current])
	}
}

func main() {
	source := "hello AND world OR alice AND NOT bob"

	scanner := NewScanner(source)

	tokens := scanner.ScanTokens()

	//fmt.Print(tokens)

	for _, token := range tokens {
		fmt.Println(token)
	}
}
