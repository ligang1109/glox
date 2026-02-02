package interpreter

import (
	"fmt"

	"github.com/ligang1109/glox/internal/stmt"
)

func (in *Interpreter) VisitExpressionStmt(exp *stmt.Expression) {
	in.evaluateExpr(exp.Exp)
}

func (in *Interpreter) VisitPrintStmt(p *stmt.Print) {
	in.evaluateExpr(p.Exp)

	fmt.Println(in.Value())
}

func (in *Interpreter) VisitVarStmt(v *stmt.Var) {
	var value any
	if v.Initializer != nil {
		in.evaluateExpr(v.Initializer)
		value = in.Value()
	}

	in.enviroment.Define(v.Variable.Lexeme, value)
}

func (in *Interpreter) VisitBlockStmt(b *stmt.Block) {
	blockEnvironment := NewEnvironment(in.enviroment)
	previous := in.enviroment
	in.enviroment = blockEnvironment

	for _, statement := range b.StatementList {
		in.evaluateStmt(statement)
	}

	in.enviroment = previous
}
