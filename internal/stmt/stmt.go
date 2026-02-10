package stmt

import (
	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/pkg/token"
)

type Statement interface {
	Name() string
	Accept(visitor StatementVisitor)
}

type StatementVisitor interface {
	VisitExpressionStmt(exp *Expression)
	VisitIfStmt(s *If)
	VisitPrintStmt(p *Print)
	VisitVarStmt(v *Var)
	VisitBlockStmt(b *Block)
	VisitWhileStmt(w *While)
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

type If struct {
	Condition  expr.Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (s *If) Name() string {
	return "If"
}

func (s *If) Accept(visitor StatementVisitor) {
	visitor.VisitIfStmt(s)
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

type Var struct {
	Variable    *token.Token
	Initializer expr.Expression
}

func (v *Var) Name() string {
	return "Var"
}

func (v *Var) Accept(visitor StatementVisitor) {
	visitor.VisitVarStmt(v)
}

type Block struct {
	StatementList []Statement
}

func (b *Block) Name() string {
	return "Block"
}

func (b *Block) Accept(visitor StatementVisitor) {
	visitor.VisitBlockStmt(b)
}

type While struct {
	Condition expr.Expression
	Body      Statement
}

func (w *While) Name() string {
	return "While"
}

func (w *While) Accept(visitor StatementVisitor) {
	visitor.VisitWhileStmt(w)
}
