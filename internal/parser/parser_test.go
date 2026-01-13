package parser

import (
	"fmt"
	"testing"

	"github.com/goinbox/gomisc"

	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/internal/scanner"
)

func parseSource(source string) {
	scanner := &scanner.Scanner{}
	tokens := scanner.Scan(source)
	for i, token := range tokens {
		fmt.Println(i, token)
	}

	parser := &Parser{}
	exp, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	content, _ := gomisc.PrettyJson(exp)
	fmt.Println(string(content))

	printer := expr.NewPrintVisitor()
	exp.Accept(printer)

	fmt.Println(printer.Graph())
}

func TestParser(t *testing.T) {
	source :=
		`
		(3+4) * -5;
		1+2;
		`

	parseSource(source)
}
