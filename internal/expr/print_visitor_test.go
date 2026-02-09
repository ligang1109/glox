package expr

import (
	"testing"

	"github.com/ligang1109/glox/pkg/token"
)

func TestPrintVisitor(t *testing.T) {
	// (3+4)*-5
	binary := &Binary{
		Left: &Grouping{
			Expression: &Binary{
				Left: &Literal{
					Value: 3,
				},
				Right: &Literal{
					Value: 4,
				},
				Operator: &token.Token{
					Lexeme: "+",
				},
			},
		},
		Right: &Unary{
			Right: &Literal{
				Value: 5,
			},
			Operator: &token.Token{
				Lexeme: "-",
			},
		},
		Operator: &token.Token{
			Lexeme: "*",
		},
	}

	visitor := NewPrintVisitor()
	binary.Accept(visitor)

	t.Log(visitor.Graph())
}

func TestVisitLogicalExpr(t *testing.T) {
	// a or b
	logical := &Logical{
		Left: &Variable{
			VarName: &token.Token{
				Lexeme: "a",
			},
		},
		Operator: &token.Token{
			Lexeme: "OR",
		},
		Right: &Variable{
			VarName: &token.Token{
				Lexeme: "b",
			},
		},
	}

	visitor := NewPrintVisitor()
	logical.Accept(visitor)

	t.Log(visitor.Graph())
}
