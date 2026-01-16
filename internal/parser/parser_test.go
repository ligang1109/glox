package parser

import (
	"fmt"
	"testing"

	"github.com/ligang1109/glox/internal/scanner"
	"github.com/ligang1109/glox/internal/stmt"
)

func parseSource(source string) {
	scanner := &scanner.Scanner{}
	tokens := scanner.Scan(source)
	for i, token := range tokens {
		fmt.Println(i, token)
	}

	parser := &Parser{}
	statementList, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, statement := range statementList {
		printer := &stmt.PrintVisitor{}
		statement.Accept(printer)

		fmt.Println(i, printer.Graph())
	}
}

func TestParser(t *testing.T) {
	source :=
		`
		(3+4) * -5;
		1+2;
		`

	parseSource(source)
}
