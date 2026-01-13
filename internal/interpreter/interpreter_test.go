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
