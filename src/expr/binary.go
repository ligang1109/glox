package expr

import "glox/token"

type Binary struct {
	Left     Expr
	Right    Expr
	Operator *token.Token
}

func (b *Binary) Name() string {
	return "Binary"
}

func (b *Binary) Accept(visitor ExprVisitor) {
	visitor.VisitBinary(b)
}
