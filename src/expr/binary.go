package expr

import "glox/token"

type Binary struct {
	Left     Expression
	Right    Expression
	Operator *token.Token
}

func (b *Binary) Name() string {
	return "Binary"
}

func (b *Binary) Accept(visitor ExpressionVisitor) {
	visitor.VisitBinary(b)
}
