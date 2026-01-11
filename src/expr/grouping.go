package expr

type Grouping struct {
	Expression Expression
}

func (g *Grouping) Name() string {
	return "Grouping"
}

func (g *Grouping) Accept(visitor ExpressionVisitor) {
	visitor.VisitGrouping(g)
}
