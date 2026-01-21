package interpreter

import (
	"fmt"
	"reflect"

	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/internal/perror"
	"github.com/ligang1109/glox/internal/stmt"
	"github.com/ligang1109/glox/pkg/token"
)

type Interpreter struct {
	value any

	enviroment *Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		enviroment: NewEnvironment(),
	}
}

func (in *Interpreter) Interpret(statementList []stmt.Statement) (err error) {
	defer func() {
		if v := recover(); v != nil {
			ok := false
			err, ok = v.(*perror.RuntimeError)
			if !ok {
				err = fmt.Errorf("Interpreter.Interpret recover from %v", v)
			}
		}
	}()

	for _, statement := range statementList {
		statement.Accept(in)
		fmt.Println(in.Value())
	}

	return nil
}

func (in *Interpreter) Value() any {
	return in.value
}

func (in *Interpreter) setValue(value any) {
	in.value = value
}

func (in *Interpreter) VisitBinaryExpr(binary *expr.Binary) {
	left := in.evaluate(binary.Left)
	right := in.evaluate(binary.Right)

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
	in.evaluate(grouping.Expression)
}

func (in *Interpreter) VisitLiteralExpr(literal *expr.Literal) {
	in.setValue(literal.Value)
}

func (in *Interpreter) VisitUnaryExpr(unary *expr.Unary) {
	value := in.evaluate(unary.Right)

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

func (in *Interpreter) evaluate(expr expr.Expression) any {
	expr.Accept(in)

	return in.Value()
}

func (in *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case bool:
		return v
	}

	return true
}

func (in *Interpreter) isEqual(v1, v2 any) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}

	return reflect.DeepEqual(v1, v2)
}

func (in *Interpreter) checkNumberOperands(operator *token.Token, values ...any) []float64 {
	result := make([]float64, len(values))
	for i, value := range values {
		v, ok := value.(float64)
		if !ok {
			in.error(operator, "Operands must be numbers.")
		}
		result[i] = v
	}

	return result
}

func (in *Interpreter) error(token *token.Token, message string) {
	panic(perror.NewRuntimeError(token, message))
}

func (in *Interpreter) VisitExpressionStmt(exp *stmt.Expression) {
	in.evaluate(exp.Exp)
}

func (in *Interpreter) VisitPrintStmt(p *stmt.Print) {
	value := in.evaluate(p.Exp)
	fmt.Println(value)
}

func (in *Interpreter) VisitVarStmt(v *stmt.Var) {
	var value any
	if v.Initializer != nil {
		value = in.evaluate(v.Initializer)
	}

	in.enviroment.Define(v.Variable.Lexeme, value)
}
