package expr

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
