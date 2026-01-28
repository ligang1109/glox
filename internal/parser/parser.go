package parser

import (
	"fmt"

	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/internal/perror"
	"github.com/ligang1109/glox/internal/stmt"
	"github.com/ligang1109/glox/pkg/log"
	"github.com/ligang1109/glox/pkg/token"
)

type Parser struct {
	tokens  []*token.Token
	current int

	hasError bool
}

func (p *Parser) HasError() bool {
	return p.hasError
}

func (p *Parser) Parse(tokens []*token.Token) []stmt.Statement {
	p.init(tokens)

	var statementList []stmt.Statement
	for {
		if p.isAtEnd() {
			break
		}

		statement, err := p.declaration()
		if err != nil {
			p.hasError = true
			log.Logger.Error(err.Error())

			p.synchronize()
		} else {
			statementList = append(statementList, statement)
		}
	}

	return statementList
}

func (p *Parser) init(tokens []*token.Token) {
	p.tokens = tokens
	p.current = 0
}

func (p *Parser) declaration() (statement stmt.Statement, err error) {
	defer func() {
		if v := recover(); v != nil {
			ok := false
			err, ok = v.(*perror.ParseError)
			if !ok {
				err = fmt.Errorf("Parser.Parse recover from %v", v)
			}
		}
	}()

	if p.match(token.Var) {
		return p.varDeclaration(), nil
	}

	return p.statement(), nil
}

func (p *Parser) varDeclaration() *stmt.Var {
	name := p.consume(token.Identifier, "Expect variable name.")

	var initializer expr.Expression
	if p.match(token.Equal) {
		initializer = p.expression()
	}

	p.consume(token.Semicolon, "Expect ';' after variable declaration.")

	return &stmt.Var{
		Variable:    name,
		Initializer: initializer,
	}
}

func (p *Parser) statement() stmt.Statement {
	if p.match(token.Print) {
		return p.printStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) printStatement() *stmt.Print {
	exp := p.expression()
	p.consume(token.Semicolon, "Expect ';' after value.")

	return &stmt.Print{
		Exp: exp,
	}
}

func (p *Parser) expressionStatement() *stmt.Expression {
	exp := p.expression()
	p.consume(token.Semicolon, "Expect ';' after expression.")

	return &stmt.Expression{
		Exp: exp,
	}
}

func (p *Parser) expression() expr.Expression {
	return p.assignment()
}

func (p *Parser) assignment() expr.Expression {
	exp := p.equality()
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

func (p *Parser) consume(tokenType token.Type, message string) *token.Token {
	if p.check(tokenType) {
		return p.advance()
	}

	p.error(nil, message)

	return nil
}

func (p *Parser) error(token *token.Token, message string) {
	if token == nil {
		token = p.peek()
	}

	panic(perror.NewParseError(token, message))
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
