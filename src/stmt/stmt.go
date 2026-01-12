package stmt

import "glox/expr"

type Statement interface {
	Name() string
	Accept(visitor StatementVisitor)
}

type StatementVisitor interface {
	VisitExpression(exp *Expression)
	VisitPrint(p *Print)
}

type Expression struct {
	exp expr.Expression
}

func (e *Expression) Name() string {
	return "Expression"
}

func (e *Expression) Accept(visitor StatementVisitor) {
	visitor.VisitExpression(e)
}

type Print struct {
	exp expr.Expression
}

func (p *Print) Name() string {
	return "Print"
}

func (p *Print) Accept(visitor StatementVisitor) {
	visitor.VisitPrint(p)
}
