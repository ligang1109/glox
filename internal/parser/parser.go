package parser

import (
	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/internal/perror"
	"github.com/ligang1109/glox/internal/stmt"
	"github.com/ligang1109/glox/pkg/token"
)

type Parser struct {
	tokens  []*token.Token
	current int
}

func (p *Parser) Parse(tokens []*token.Token) (statementList []stmt.Statement, err *perror.ParseError) {
	p.init(tokens)

	defer func() {
		if v := recover(); v != nil {
			err = v.(*perror.ParseError)
		}
	}()

	for {
		if p.isAtEnd() {
			break
		}

		statementList = append(statementList, p.statement())
	}

	return statementList, nil
}

func (p *Parser) init(tokens []*token.Token) {
	p.tokens = tokens
	p.current = 0
}

func (p *Parser) statement() stmt.Statement {
	if p.match(token.Print) {
		return p.printStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) printStatement() *stmt.Print {
	exp := p.expression()
	if !p.match(token.Semicolon) {
		p.error("Expect ';' after value.")
	}

	return &stmt.Print{
		Exp: exp,
	}
}

func (p *Parser) expressionStatement() *stmt.Expression {
	exp := p.expression()
	if !p.match(token.Semicolon) {
		p.error("Expect ';' after expression.")
	}

	return &stmt.Expression{
		Exp: exp,
	}
}

func (p *Parser) expression() expr.Expression {
	return p.equality()
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

	return p.primary()
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
		if p.match(token.RightParen) {
			return &expr.Grouping{
				Expression: exp,
			}
		}
		p.error("Expect ')' after expression.")
	}

	p.error("Unexpected token.")

	return nil
}

func (p *Parser) match(tokenTypes ...token.Type) bool {
	for _, tt := range tokenTypes {
		if p.check(tt) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType token.Type) bool {
	token := p.peek()
	if token == nil {
		return false
	}

	return token.Type == tokenType
}

func (p *Parser) isAtEnd() bool {
	if p.current >= len(p.tokens)-1 {
		return true
	}

	return false
}

func (p *Parser) peek() *token.Token {
	if p.isAtEnd() {
		return nil
	}

	return p.tokens[p.current]
}

func (p *Parser) advance() *token.Token {
	token := p.peek()
	if token == nil {
		return nil
	}

	p.current++

	return token
}

func (p *Parser) previous() *token.Token {
	return p.tokens[(p.current - 1)]
}

func (p *Parser) error(message string) {
	panic(perror.NewParseError(p.peek(), message))
}

func (p *Parser) synchronize() {
	previous := p.advance()
	if previous == nil || previous.Type == token.Semicolon {
		return
	}

	for {
		current := p.peek()
		if current == nil {
			return
		}

		switch current.Type {
		case token.Semicolon, token.Class, token.Fun, token.Var, token.For, token.If, token.While, token.Print, token.Return:
			return
		}

		p.advance()
	}
}
