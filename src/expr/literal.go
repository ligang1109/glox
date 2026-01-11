package expr

type Literal struct {
	Value any
}

func (l *Literal) Name() string {
	return "Literal"
}

func (l *Literal) Accept(visitor ExpressionVisitor) {
	visitor.VisitLiteral(l)
}
