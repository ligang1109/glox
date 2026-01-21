package interpreter

import (
	"fmt"
	"testing"

	"github.com/ligang1109/glox/internal/parser"
	"github.com/ligang1109/glox/internal/scanner"
)

func interpret(source string) {
	scanner := &scanner.Scanner{}
	tokens := scanner.Scan(source)
	for i, token := range tokens {
		fmt.Println(i, token)
	}
	parser := &parser.Parser{}
	statementList := parser.Parse(tokens)
	if parser.HasError() {
		return
	}

	interpreter := &Interpreter{}
	e := interpreter.Interpret(statementList)
	if e != nil {
		fmt.Println("interpreter.Interpret error:", e)
		return
	}
}

func TestInterpreter(t *testing.T) {
	interpret(`
		3-1;
		(3+4) * -5;
		print true;
		`)
}
