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
	t.Log(graph)
}
