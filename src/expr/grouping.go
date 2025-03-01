package expr

type Grouping struct {
	Expression Expr
}

func (g *Grouping) Name() string {
	return "Grouping"
}
