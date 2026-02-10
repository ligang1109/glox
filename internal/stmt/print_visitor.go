package stmt

import (
	"github.com/ligang1109/glox/internal/expr"
)

type PrintVisitor struct {
	graph string
}

func (pv *PrintVisitor) Graph() string {
	return pv.graph
}

func (pv *PrintVisitor) drawExpression(exp expr.Expression) {
	exprPrinter := expr.NewPrintVisitor()
	exp.Accept(exprPrinter)

	pv.graph = exprPrinter.Graph()
}

func (pv *PrintVisitor) VisitExpressionStmt(exp *Expression) {
	pv.drawExpression(exp.Exp)
}

func (pv *PrintVisitor) VisitIfStmt(s *If) {
	pv.drawExpression(s.Condition)
	s.ThenBranch.Accept(pv)
	if s.ElseBranch != nil {
		s.ElseBranch.Accept(pv)
	}
}

func (pv *PrintVisitor) VisitPrintStmt(p *Print) {
	pv.drawExpression(p.Exp)
}

func (pv *PrintVisitor) VisitVarStmt(v *Var) {
	pv.drawExpression(&expr.Variable{
		VarName: v.Variable,
	})
}

func (pv *PrintVisitor) VisitBlockStmt(b *Block) {
	for _, statement := range b.StatementList {
		statement.Accept(pv)
	}
}

func (pv *PrintVisitor) VisitWhileStmt(w *While) {
	pv.drawExpression(w.Condition)
	w.Body.Accept(pv)
}
