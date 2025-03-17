package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

const src string = `package foo

import (
	"fmt"
	"time"
)

func baz() {
	fmt.Println("Hello, world!")
}

type A int

const b = "testing"

func bar() {
	fmt.Println(time.Now())
}`

type ByName []*dst.FuncDecl

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name.Name < a[j].Name.Name }

// Moves all top-level functions to the end, sorted in alphabetical order.
// The "source file" is given as a string (rather than e.g. a filename).
func SortFunctions(src string) (string, error) {
	d := decorator.NewDecorator(nil)
	f, err := d.Parse(src)
	if err != nil {
		return "", err
	}

	var functions []*dst.FuncDecl
	var otherDecls []dst.Decl

	for _, decl := range f.Decls {
		if fn, ok := decl.(*dst.FuncDecl); ok {
			functions = append(functions, fn)
		} else {
			otherDecls = append(otherDecls, decl)
		}
	}

	sort.Sort(ByName(functions))

	f.Decls = otherDecls
	for _, fn := range functions {
		f.Decls = append(f.Decls, fn)
	}

	var buf bytes.Buffer
	err = decorator.Fprint(&buf, f)
	if err != nil {
		return "", err
	}

	return buf.String(), nil

}

func main() {
	f, err := decorator.Parse(src)
	if err != nil {
		log.Fatal(err)
	}

	// Print AST
	err = dst.Fprint(os.Stdout, f, nil)
	if err != nil {
		log.Fatal(err)
	}

	sorted, err := SortFunctions(src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", sorted)
}
