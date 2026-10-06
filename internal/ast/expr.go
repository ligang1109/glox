package ast

import "github.com/ligang1109/glox/internal/token"

type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (b *Binary) Name() string {
	return "Binary"
}

func (b *Binary) Accept(visitor ExprVisitor) Value {
	return visitor.VisitBinary(b)
}

type Literal struct {
	Value Value
}

func (l *Literal) Name() string {
	return "Literal"
}

func (l *Literal) Accept(visitor ExprVisitor) Value {
	return visitor.VisitLiteral(l)
}

type Unary struct {
	Right    Expr
	Operator *token.Token
}

func (u *Unary) Name() string {
	return "Unary"
}

func (u *Unary) Accept(visitor ExprVisitor) Value {
	return visitor.VisitUnary(u)
}

type Grouping struct {
	Expr Expr
}

func (g *Grouping) Name() string {
	return "Grouping"
}

func (g *Grouping) Accept(visitor ExprVisitor) Value {
	return visitor.VisitGrouping(g)
}
