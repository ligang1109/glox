package parser

import (
	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/pkg/token"
)

func (p *Parser) expression() expr.Expression {
	return p.assignment()
}

func (p *Parser) assignment() expr.Expression {
	exp := p.logicOr()
	if !p.match(token.Equal) {
		return exp
	}

	equal := p.previous()
	value := p.assignment()
	variable, ok := exp.(*expr.Variable)
	if !ok {
		p.error(equal, "Invalid assignment target.")
	}

	return &expr.Assign{
		VarName: variable.VarName,
		Value:   value,
	}
}

func (p *Parser) logicOr() expr.Expression {
	exp := p.logicAnd()
	for {
		if p.match(token.Or) {
			operator := p.previous()
			right := p.logicAnd()
			exp = &expr.Logical{
				Left:     exp,
				Operator: operator,
				Right:    right,
			}
		} else {
			break
		}
	}

	return exp
}

func (p *Parser) logicAnd() expr.Expression {
	exp := p.equality()
	for {
		if p.match(token.And) {
			operator := p.previous()
			right := p.equality()
			exp = &expr.Logical{
				Left:     exp,
				Operator: operator,
				Right:    right,
			}
		} else {
			break
		}
	}

	return exp
}

func (p *Parser) equality() expr.Expression {
	exp := p.comparison()
	for {
		if p.match(token.BangEqual, token.EqualEqual) {
			operator := p.previous()
			right := p.comparison()
			exp = &expr.Binary{
				Left:     exp,
				Right:    right,
				Operator: operator,
			}
		} else {
			break
		}
	}

	return exp
}

func (p *Parser) comparison() expr.Expression {
	exp := p.term()
	for {
		if p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
			operator := p.previous()
			right := p.term()
			exp = &expr.Binary{
				Left:     exp,
				Right:    right,
				Operator: operator,
			}
		} else {
			break
		}
	}

	return exp
}

func (p *Parser) term() expr.Expression {
	exp := p.factor()
	for {
		if p.match(token.Minus, token.Plus) {
			operator := p.previous()
			right := p.factor()
			exp = &expr.Binary{
				Left:     exp,
				Right:    right,
				Operator: operator,
			}
		} else {
			break
		}
	}

	return exp
}

func (p *Parser) factor() expr.Expression {
	exp := p.unary()
	for {
		if p.match(token.Slash, token.Star) {
			operator := p.previous()
			right := p.unary()
			exp = &expr.Binary{
				Left:     exp,
				Right:    right,
				Operator: operator,
			}
		} else {
			break
		}
	}

	return exp
}

func (p *Parser) unary() expr.Expression {
	if p.match(token.Bang, token.Minus) {
		operator := p.previous()
		right := p.unary()
		return &expr.Unary{
			Right:    right,
			Operator: operator,
		}
	}

	return p.call()
}

func (p *Parser) call() expr.Expression {
	exp := p.primary()
	for {
		if p.match(token.LeftParen) {
			exp = p.finishCall(exp)
		} else {
			break
		}
	}

	return exp
}

func (p *Parser) finishCall(callee expr.Expression) *expr.Call {
	arguments := []expr.Expression{}
	if !p.check(token.RightParen) {
		for {
			arguments = append(arguments, p.expression())
			if !p.match(token.Comma) {
				break
			}
		}
	}

	paren := p.consume(token.RightParen, "Expect ')' after arguments.")
	return &expr.Call{
		Callee:    callee,
		Paren:     paren,
		Arguments: arguments,
	}
}

func (p *Parser) primary() expr.Expression {
	if p.match(token.Number, token.String) {
		return &expr.Literal{
			Value: p.previous().Literal,
		}
	}

	if p.match(token.True) {
		return &expr.Literal{
			Value: true,
		}
	}
	if p.match(token.False) {
		return &expr.Literal{
			Value: false,
		}
	}
	if p.match(token.Nil) {
		return &expr.Literal{
			Value: nil,
		}
	}

	if p.match(token.LeftParen) {
		exp := p.expression()
		p.consume(token.RightParen, "Expect ')' after expression.")

		return &expr.Grouping{
			Expression: exp,
		}
	}

	if p.match(token.Identifier) {
		return &expr.Variable{
			VarName: p.previous(),
		}
	}

	p.error(nil, "Unexpected token.")

	return nil
}
