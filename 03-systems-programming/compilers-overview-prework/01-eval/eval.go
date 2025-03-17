package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

// Given an expression containing only int types, evaluate
// the expression and return the result.
func Evaluate(expr ast.Expr) (int, error) {
	switch e := expr.(type) {
	case *ast.BinaryExpr:
		left, err := Evaluate(e.X)
		if err != nil {
			return 0, err
		}
		right, err := Evaluate(e.Y)
		if err != nil {
			return 0, err
		}

		switch e.Op {
		case token.ADD:
			return left + right, nil
		case token.SUB:
			return left - right, nil
		case token.MUL:
			return left * right, nil
		case token.QUO:
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return left / right, nil

		}
	case *ast.ParenExpr:
		return Evaluate(e.X)
	case *ast.BasicLit:
		return strconv.Atoi(e.Value)
	default:
		return 0, fmt.Errorf("unsupported expr type: %T", e)
	}
	return 0, nil
}

func main() {
	expr, err := parser.ParseExpr("1 + 2 - 3 * 4")
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	err = ast.Print(fset, expr)
	if err != nil {
		log.Fatal(err)
	}
}
