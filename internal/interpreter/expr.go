package interpreter

import (
	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/pkg/token"
)

func (in *Interpreter) VisitBinaryExpr(binary *expr.Binary) {
	in.evaluateExpr(binary.Left)
	left := in.Value()

	in.evaluateExpr(binary.Right)
	right := in.Value()

	switch binary.Operator.Type {
	case token.Minus:
		values := in.checkNumberOperands(binary.Operator, left, right)
		in.setValue(values[0] - values[1])
		return
	case token.Slash:
		values := in.checkNumberOperands(binary.Operator, left, right)
		in.setValue(values[0] / values[1])
		return
	case token.Star:
		values := in.checkNumberOperands(binary.Operator, left, right)
		in.setValue(values[0] * values[1])
		return
	case token.Greater:
		values := in.checkNumberOperands(binary.Operator, left, right)
		in.setValue(values[0] > values[1])
		return
	case token.GreaterEqual:
		values := in.checkNumberOperands(binary.Operator, left, right)
		in.setValue(values[0] >= values[1])
		return
	case token.Less:
		values := in.checkNumberOperands(binary.Operator, left, right)
		in.setValue(values[0] < values[1])
		return
	case token.LessEqual:
		values := in.checkNumberOperands(binary.Operator, left, right)
		in.setValue(values[0] <= values[1])
		return
	case token.BangEqual:
		in.setValue(!in.isEqual(left, right))
		return
	case token.EqualEqual:
		in.setValue(in.isEqual(left, right))
		return
	case token.Plus:
		lv, ok := left.(float64)
		if ok {
			rv, ok := right.(float64)
			if ok {
				in.setValue(lv + rv)
				return
			}
			in.error(binary.Operator, "Right is not number.")
		}

		ls, ok := left.(string)
		if ok {
			rs, ok := right.(string)
			if ok {
				in.setValue(ls + rs)
				return
			}
			in.error(binary.Operator, "Right is not string.")
		}

		in.error(binary.Operator, "Operands must be two numbers or two strings.")
	}

	in.error(binary.Operator, "Unreachable.")
}

func (in *Interpreter) VisitGroupingExpr(grouping *expr.Grouping) {
	in.evaluateExpr(grouping.Expression)
}

func (in *Interpreter) VisitLiteralExpr(literal *expr.Literal) {
	in.setValue(literal.Value)
}

func (in *Interpreter) VisitUnaryExpr(unary *expr.Unary) {
	in.evaluateExpr(unary.Right)
	value := in.Value()

	switch unary.Operator.Type {
	case token.Minus:
		in.setValue(-(value.(float64)))
		return
	case token.Bang:
		in.setValue(!in.isTruthy(value))
		return
	}

	in.error(unary.Operator, "Unreachable.")
}

func (in *Interpreter) VisitVariableExpr(variable *expr.Variable) {
	in.setValue(in.enviroment.Value(variable.VarName))
}

func (in *Interpreter) VisitAssignExpr(assign *expr.Assign) {
	in.evaluateExpr(assign.Value)
	in.enviroment.Assign(assign.VarName, in.Value())
}

func (in *Interpreter) VisitLogicalExpr(logical *expr.Logical) {
	in.evaluateExpr(logical.Left)

	if in.isTruthy(in.Value()) {
		if logical.Operator.Type == token.Or {
			return
		}
	} else {
		if logical.Operator.Type == token.And {
			return
		}
	}

	in.evaluateExpr(logical.Right)
}
