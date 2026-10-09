package printer_test

import (
	"testing"

	"github.com/ligang1109/glox/internal/ast"
	"github.com/ligang1109/glox/internal/printer"
	"github.com/ligang1109/glox/internal/token"
)

func TestPrintExpr(t *testing.T) {
	// (3+4)*-5
	expr := &ast.Binary{
		Left: &ast.Grouping{
			Expr: &ast.Binary{
				Left: &ast.Literal{
					Value: 3,
				},
				Operator: &token.Token{
					Lexeme: "+",
				},
				Right: &ast.Literal{
					Value: 4,
				},
			},
		},
		Operator: &token.Token{
			Lexeme: "*",
		},
		Right: &ast.Unary{
			Right: &ast.Literal{
				Value: 5,
			},
			Operator: &token.Token{
				Lexeme: "-",
			},
		},
	}
	p := &printer.Printer{}
	graph := p.PrintExpr(expr)

	want := `flowchart TD
n0["Binary *"]
n1["Grouping"]
n0 --> n1
n2["Binary +"]
n1 --> n2
n3["Literal 3"]
n2 --> n3
n4["Literal 4"]
n2 --> n4
n5["Unary -"]
n0 --> n5
n6["Literal 5"]
n5 --> n6
`
	if graph != want {
		t.Errorf("got:\n%s\nwant:\n%s", graph, want)
	}
}

func TestPrintExprStringLiteral(t *testing.T) {
	// he said "hi"
	expr := &ast.Literal{
		Value: `he said "hi"`,
	}

	p := &printer.Printer{}
	graph := p.PrintExpr(expr)

	want := `flowchart TD
n0["Literal #quot;he said \#quot;hi\#quot;#quot;"]
`
	if graph != want {
		t.Errorf("got:\n%s\nwant:\n%s", graph, want)
	}
}

func TestPrintExprReuse(t *testing.T) {
	expr := &ast.Binary{
		Left: &ast.Literal{
			Value: 1,
		},
		Operator: &token.Token{
			Lexeme: "+",
		},
		Right: &ast.Literal{
			Value: 2,
		},
	}

	p := &printer.Printer{}

	first := p.PrintExpr(expr)
	second := p.PrintExpr(expr)

	if first != second {
		t.Errorf("printing the same tree twice differs:\nfirst:\n%s\nsecond:\n%s", first, second)
	}
}
