package interpreter

import (
	"fmt"
	"testing"

	"glox/parser"
	"glox/scanner"
)

func interpret(source string) {
	scanner := &scanner.Scanner{}
	tokens := scanner.Scan(source)
	for i, token := range tokens {
		fmt.Println(i, token)
	}
	parser := &parser.Parser{}
	exp, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println("parser.Parse error:", err)
		return
	}

	interpreter := &Interpreter{}
	interpreter.Interpret(exp)
}

func TestInterpreter(t *testing.T) {
	interpret(`
		// 3-"abc";
		(3+4) * -5;
		`)
}
