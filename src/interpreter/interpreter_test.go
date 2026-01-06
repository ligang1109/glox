package interpreter

import (
	"testing"

	"glox/parser"
	"glox/scanner"
)

func interpret(source string) any {
	scanner := &scanner.Scanner{}
	tokens := scanner.Scan(source)
	parser := &parser.Parser{}
	exp, _ := parser.Parse(tokens)

	interpreter := NewInterpreter()
	exp.Accept(interpreter)

	return interpreter.Value()
}

func TestInterpreter(t *testing.T) {
	v := interpret(`
		(3+4) * -5;
		`)

	t.Log(v)
}
