package ast

import "github.com/ligang1109/glox/internal/token"

type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (b *Binary) Accept(visitor ExprVisitor) VisitResult {
	return visitor.VisitBinary(b)
}

type Literal struct {
	Value Value
}

func (l *Literal) Accept(visitor ExprVisitor) VisitResult {
	return visitor.VisitLiteral(l)
}

type Unary struct {
	Right    Expr
	Operator *token.Token
}

func (u *Unary) Accept(visitor ExprVisitor) VisitResult {
	return visitor.VisitUnary(u)
}

type Grouping struct {
	Expr Expr
}

func (g *Grouping) Accept(visitor ExprVisitor) VisitResult {
	return visitor.VisitGrouping(g)
}
