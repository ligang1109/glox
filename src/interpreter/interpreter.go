package interpreter

import (
	"fmt"
	"reflect"

	"glox/expr"
	"glox/perror"
	"glox/token"
)

type Interpreter struct {
	value any
}

func (p *Interpreter) Interpret(exp expr.Expr) {
	defer func() {
		if v := recover(); v != nil {
			err := v.(*perror.RuntimeError)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	p.evaluate(exp)
	fmt.Println(p.Value())
}

func (p *Interpreter) Value() any {
	return p.value
}

func (p *Interpreter) setValue(value any) {
	p.value = value
}

func (p *Interpreter) VisitBinary(binary *expr.Binary) {
	p.evaluate(binary.Left)
	left := p.Value()

	p.evaluate(binary.Right)
	right := p.Value()

	switch binary.Operator.Type {
	case token.Minus:
		values := p.checkNumberOperands(binary.Operator, left, right)
		p.setValue(values[0] - values[1])
		return
	case token.Slash:
		values := p.checkNumberOperands(binary.Operator, left, right)
		p.setValue(values[0] / values[1])
		return
	case token.Star:
		values := p.checkNumberOperands(binary.Operator, left, right)
		p.setValue(values[0] * values[1])
		return
	case token.Greater:
		values := p.checkNumberOperands(binary.Operator, left, right)
		p.setValue(values[0] > values[1])
		return
	case token.GreaterEqual:
		values := p.checkNumberOperands(binary.Operator, left, right)
		p.setValue(values[0] >= values[1])
		return
	case token.Less:
		values := p.checkNumberOperands(binary.Operator, left, right)
		p.setValue(values[0] < values[1])
		return
	case token.LessEqual:
		values := p.checkNumberOperands(binary.Operator, left, right)
		p.setValue(values[0] <= values[1])
		return
	case token.BangEqual:
		p.setValue(!p.isEqual(left, right))
		return
	case token.EqualEqual:
		p.setValue(p.isEqual(left, right))
		return
	case token.Plus:
		lv, ok := left.(float64)
		if ok {
			rv, ok := right.(float64)
			if ok {
				p.setValue(lv + rv)
				return
			}
			p.error(binary.Operator, "Right is not number.")
		}

		ls, ok := left.(string)
		if ok {
			rs, ok := right.(string)
			if ok {
				p.setValue(ls + rs)
				return
			}
			p.error(binary.Operator, "Right is not string.")
		}

		p.error(binary.Operator, "Operands must be two numbers or two strings.")
	}

	p.error(binary.Operator, "Unreachable.")
}

func (p *Interpreter) VisitGrouping(grouping *expr.Grouping) {
	p.evaluate(grouping.Expression)
}

func (p *Interpreter) VisitLiteral(literal *expr.Literal) {
	p.setValue(literal.Value)
}

func (p *Interpreter) VisitUnary(unary *expr.Unary) {
	p.evaluate(unary.Right)

	switch unary.Operator.Type {
	case token.Minus:
		v := p.Value().(float64)
		p.setValue(-v)
		return
	case token.Bang:
		v := p.Value()
		p.setValue(!p.isTruthy(v))
		return
	}

	p.error(unary.Operator, "Unreachable.")
}

func (p *Interpreter) evaluate(expr expr.Expr) {
	expr.Accept(p)
}

func (p *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case bool:
		return v
	}

	return true
}

func (p *Interpreter) isEqual(v1, v2 any) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}

	return reflect.DeepEqual(v1, v2)
}

func (p *Interpreter) checkNumberOperands(operator *token.Token, values ...any) []float64 {
	result := make([]float64, len(values))
	for i, value := range values {
		v, ok := value.(float64)
		if !ok {
			p.error(operator, "Operands must be numbers.")
		}
		result[i] = v
	}

	return result
}

func (p *Interpreter) error(token *token.Token, message string) {
	panic(perror.NewRuntimeError(token, message))
}
