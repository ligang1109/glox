package expr

type Expr interface {
	Name() string
	Accept(visitor ExprVisitor)
}

type ExprVisitor interface {
	VisitBinary(binary *Binary)
	VisitGrouping(grouping *Grouping)
	VisitLiteral(literal *Literal)
	VisitUnary(unary *Unary)
}
