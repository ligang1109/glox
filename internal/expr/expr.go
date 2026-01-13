package expr

import "github.com/ligang1109/glox/pkg/token"

type Expression interface {
	Name() string
	Accept(visitor ExpressionVisitor)
}

type ExpressionVisitor interface {
	VisitBinary(binary *Binary)
	VisitGrouping(grouping *Grouping)
	VisitLiteral(literal *Literal)
	VisitUnary(unary *Unary)
}

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

type Grouping struct {
	Expression Expression
}

func (g *Grouping) Name() string {
	return "Grouping"
}

func (g *Grouping) Accept(visitor ExpressionVisitor) {
	visitor.VisitGrouping(g)
}

type Literal struct {
	Value any
}

func (l *Literal) Name() string {
	return "Literal"
}

func (l *Literal) Accept(visitor ExpressionVisitor) {
	visitor.VisitLiteral(l)
}

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
