package interpreter

import (
	"reflect"

	"github.com/goinbox/ds"

	"glox/expr"
	"glox/token"
)

type Interpreter struct {
	values ds.Stack[any]
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		values: &ds.SimpleStack[any]{},
	}
}

func (p *Interpreter) Value() any {
	v, _ := p.values.Pop()

	return v
}

func (p *Interpreter) VisitBinary(binary *expr.Binary) {
	p.evaluate(binary.Left)
	left := p.Value()

	p.evaluate(binary.Right)
	right := p.Value()

	switch binary.Operator.Type {
	case token.Minus:
		p.values.Push(left.(float64) - right.(float64))
		return
	case token.Slash:
		p.values.Push(left.(float64) / right.(float64))
		return
	case token.Star:
		p.values.Push(left.(float64) * right.(float64))
		return
	case token.Greater:
		p.values.Push(left.(float64) > right.(float64))
		return
	case token.GreaterEqual:
		p.values.Push(left.(float64) >= right.(float64))
		return
	case token.Less:
		p.values.Push(left.(float64) < right.(float64))
		return
	case token.LessEqual:
		p.values.Push(left.(float64) <= right.(float64))
		return
	case token.BangEqual:
		p.values.Push(!p.isEqual(left, right))
		return
	case token.EqualEqual:
		p.values.Push(p.isEqual(left, right))
		return
	case token.Plus:
		lv, ok := left.(float64)
		if ok {
			rv, ok := right.(float64)
			if ok {
				p.values.Push(lv + rv)
				return
			}
			panic("right not number")
		}

		ls, ok := left.(string)
		if ok {
			rs, ok := right.(string)
			if ok {
				p.values.Push(ls + rs)
				return
			}
			panic("right not string")
		}

		panic("left error")
	}

	panic("Unreachable.")
}

func (p *Interpreter) VisitGrouping(grouping *expr.Grouping) {
	p.evaluate(grouping.Expression)
}

func (p *Interpreter) VisitLiteral(literal *expr.Literal) {
	p.values.Push(literal.Value)
}

func (p *Interpreter) VisitUnary(unary *expr.Unary) {
	p.evaluate(unary.Right)

	switch unary.Operator.Type {
	case token.Minus:
		v := p.Value().(float64)
		p.values.Push(-v)
		return
	case token.Bang:
		v := p.Value()
		p.values.Push(!p.isTruthy(v))
		return
	}

	panic("Unreachable.")
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
