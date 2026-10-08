package printer

import (
	"fmt"

	"github.com/ligang1109/glox/internal/ast"
)

func (p *Printer) VisitBinary(binary *ast.Binary) ast.VisitResult {
	return &node{
		label: fmt.Sprintf("Binary %s", binary.Operator.Lexeme),
		children: []*node{
			p.exprToNode(binary.Left),
			p.exprToNode(binary.Right),
		},
	}
}

func (p *Printer) VisitGrouping(grouping *ast.Grouping) ast.VisitResult {
	return &node{
		label: "Grouping",
		children: []*node{
			p.exprToNode(grouping.Expr),
		},
	}
}

func (p *Printer) VisitLiteral(literal *ast.Literal) ast.VisitResult {
	if s, ok := literal.Value.(string); ok {
		return &node{
			label:    fmt.Sprintf("Literal %q", s),
			children: []*node{},
		}
	}

	return &node{
		label:    fmt.Sprintf("Literal %v", literal.Value),
		children: []*node{},
	}
}

func (p *Printer) VisitUnary(unary *ast.Unary) ast.VisitResult {
	return &node{
		label: fmt.Sprintf("Unary %s", unary.Operator.Lexeme),
		children: []*node{
			p.exprToNode(unary.Right),
		},
	}
}
