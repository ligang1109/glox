package expr

type Literal struct {
	Value any
}

func (l *Literal) Name() string {
	return "Literal"
}

func (l *Literal) Accept(visitor ExprVisitor) {
	visitor.VisitLiteral(l)
}
