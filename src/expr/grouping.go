package expr

type Grouping struct {
	Expression Expr
}

func (g *Grouping) Name() string {
	return "Grouping"
}

func (g *Grouping) Accept(visitor ExprVisitor) {
	visitor.VisitGrouping(g)
}
