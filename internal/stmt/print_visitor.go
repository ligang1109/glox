package stmt

import "github.com/ligang1109/glox/internal/expr"

type PrintVisitor struct {
	graph string
}

func (v *PrintVisitor) Graph() string {
	return v.graph
}

func (v *PrintVisitor) drawExpression(exp expr.Expression) {
	exprPrinter := expr.NewPrintVisitor()
	exp.Accept(exprPrinter)

	v.graph = exprPrinter.Graph()
}

func (v *PrintVisitor) VisitExpression(exp *Expression) {
	v.drawExpression(exp.Exp)
}

func (v *PrintVisitor) VisitPrint(p *Print) {
	v.drawExpression(p.Exp)
}
