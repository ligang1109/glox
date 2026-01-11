package expr

import "glox/token"

type Unary struct {
	Right    Expression
	Operator *token.Token
}

func (u *Unary) Name() string {
	return "Unary"
}

func (u *Unary) Accept(visitor ExpressionVisitor) {
	visitor.VisitUnary(u)
}
