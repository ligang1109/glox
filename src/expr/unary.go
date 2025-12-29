package expr

import "glox/token"

type Unary struct {
	Right    Expr
	Operator *token.Token
}

func (u *Unary) Name() string {
	return "Unary"
}

func (u *Unary) Accept(visitor ExprVisitor) {
	visitor.VisitUnary(u)
}
