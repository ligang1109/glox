package expr

import "github.com/ligang1109/glox/pkg/token"

type Expression interface {
	Name() string
	Accept(visitor ExpressionVisitor)
}

type ExpressionVisitor interface {
	VisitBinaryExpr(binary *Binary)
	VisitGroupingExpr(grouping *Grouping)
	VisitLiteralExpr(literal *Literal)
	VisitUnaryExpr(unary *Unary)
	VisitVariableExpr(variable *Variable)
	VisitAssignExpr(assign *Assign)
	VisitLogicalExpr(logical *Logical)
}

type Binary struct {
	Left     Expression
	Operator *token.Token
	Right    Expression
}

func (b *Binary) Name() string {
	return "Binary"
}

func (b *Binary) Accept(visitor ExpressionVisitor) {
	visitor.VisitBinaryExpr(b)
}

type Grouping struct {
	Expression Expression
}

func (g *Grouping) Name() string {
	return "Grouping"
}

func (g *Grouping) Accept(visitor ExpressionVisitor) {
	visitor.VisitGroupingExpr(g)
}

type Literal struct {
	Value any
}

func (l *Literal) Name() string {
	return "Literal"
}

func (l *Literal) Accept(visitor ExpressionVisitor) {
	visitor.VisitLiteralExpr(l)
}

type Unary struct {
	Right    Expression
	Operator *token.Token
}

func (u *Unary) Name() string {
	return "Unary"
}

func (u *Unary) Accept(visitor ExpressionVisitor) {
	visitor.VisitUnaryExpr(u)
}

type Variable struct {
	VarName *token.Token
}

func (v *Variable) Name() string {
	return "Variable"
}

func (v *Variable) Accept(visitor ExpressionVisitor) {
	visitor.VisitVariableExpr(v)
}

type Assign struct {
	VarName *token.Token
	Value   Expression
}

func (a *Assign) Name() string {
	return "Assign"
}

func (a *Assign) Accept(visitor ExpressionVisitor) {
	visitor.VisitAssignExpr(a)
}

type Logical struct {
	Left     Expression
	Operator *token.Token
	Right    Expression
}

func (l *Logical) Name() string {
	return "Logical"
}

func (l *Logical) Accept(visitor ExpressionVisitor) {
	visitor.VisitLogicalExpr(l)
}
