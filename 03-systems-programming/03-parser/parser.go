package main

import (
	"compiler/scanner"
	"fmt"
	"strings"
)

type Expr interface {
	String() string
}

type BinaryExpr struct {
	Left     Expr
	Operator scanner.Token
	Right    Expr
}

type UnaryExpr struct {
	Operator scanner.Token
	Right    Expr
}

type LiteralExpr struct {
	Value scanner.Token
}

func (b *BinaryExpr) String() string {
	return fmt.Sprintf("%s(%v, %v)", strings.ToUpper(b.Operator.Lexeme), b.Left, b.Right)
}

func (u *UnaryExpr) String() string {
	return fmt.Sprintf("%s(%v)", strings.ToUpper(u.Operator.Lexeme), u.Right)
}

func (l *LiteralExpr) String() string {
	return fmt.Sprintf("TERM(%s)", l.Value.Lexeme)
}

type Parser struct {
	tokens  []scanner.Token
	current int
}

func NewParser(tokens []scanner.Token) *Parser {
	return &Parser{tokens: tokens}
}

// util parser fns
func (p *Parser) match(types ...scanner.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == tokenType
}

func (p *Parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens)
}

func (p *Parser) peek() scanner.Token {
	if p.isAtEnd() {
		return scanner.Token{Type: scanner.EOF, Lexeme: "", Literal: nil}
	}
	return p.tokens[p.current]
}

func (p *Parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) parseQuery() Expr {
	return p.parseOr()
}

func (p *Parser) parseOr() Expr {
	expr := p.parseAnd()

	for p.match(scanner.OR) {
		operator := p.previous()
		right := p.parseAnd()
		expr = &BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) parseAnd() Expr {
	expr := p.parseNot()

	for p.match(scanner.AND) {
		operator := p.previous()
		right := p.parseNot()
		expr = &BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr
}

func (p *Parser) parseNot() Expr {
	if p.match(scanner.NOT) {
		operator := p.previous()
		right := p.parsePrimary()
		return &UnaryExpr{Operator: operator, Right: right}
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() Expr {
	if p.match(scanner.IDENTIFIER) {
		return &LiteralExpr{Value: p.previous()}
	}

	if p.match(scanner.LEFT_PAREN) {
		expr := p.parseQuery()
		if !p.match(scanner.RIGHT_PAREN) {
			panic("Expected ')' after expression. ")
		}
		return expr
	}

	panic("Expected expression")
}

func main() {
	source := "hello AND world OR alice AND NOT bob"

	scanner := scanner.NewScanner(source)

	tokens := scanner.ScanTokens()
	fmt.Println(tokens)

	parser := NewParser(tokens)
	expression := parser.parseQuery()

	fmt.Println(expression)
}
