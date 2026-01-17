package stmt

import "github.com/ligang1109/glox/internal/expr"

type Statement interface {
	Name() string
	Accept(visitor StatementVisitor)
}

type StatementVisitor interface {
	VisitExpressionStmt(exp *Expression)
	VisitPrintStmt(p *Print)
}

type Expression struct {
	Exp expr.Expression
}

func (e *Expression) Name() string {
	return "Expression"
}

func (e *Expression) Accept(visitor StatementVisitor) {
	visitor.VisitExpressionStmt(e)
}

type Print struct {
	Exp expr.Expression
}

func (p *Print) Name() string {
	return "Print"
}

func (p *Print) Accept(visitor StatementVisitor) {
	visitor.VisitPrintStmt(p)
}
