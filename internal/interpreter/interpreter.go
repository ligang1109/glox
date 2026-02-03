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
		enviroment: NewEnvironment(nil),
	}
}

func (in *Interpreter) Interpret(statementList []stmt.Statement) error {
	for _, statement := range statementList {
		err := in.evaluateStmt(statement)
		if err != nil {
			return err
		}

		// fmt.Println(in.Value())
	}

	return nil
}

func (in *Interpreter) Value() any {
	return in.value
}

func (in *Interpreter) setValue(value any) {
	in.value = value
}

func (in *Interpreter) evaluate(name string, f func()) (err error) {
	defer func() {
		if v := recover(); v != nil {
			ok := false
			err, ok = v.(*perror.RuntimeError)
			if !ok {
				err = fmt.Errorf("%s recover from %v", name, v)
			}
		}
	}()

	f()

	return nil
}

func (in *Interpreter) evaluateStmt(statement stmt.Statement) (err error) {
	return in.evaluate("evaluateStmt",
		func() {
			statement.Accept(in)
		})
}

func (in *Interpreter) evaluateExpr(expression expr.Expression) (err error) {
	return in.evaluate("evaluateExpr",
		func() {
			expression.Accept(in)
		})
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
