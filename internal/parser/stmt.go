package parser

import (
	"github.com/ligang1109/glox/internal/expr"
	"github.com/ligang1109/glox/internal/stmt"
	"github.com/ligang1109/glox/pkg/token"
)

func (p *Parser) declaration() stmt.Statement {
	if p.match(token.Var) {
		return p.varDeclaration()
	}

	return p.statement()
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
	if p.match(token.If) {
		return p.ifStatement()
	}

	if p.match(token.Print) {
		return p.printStatement()
	}

	if p.match(token.LeftBrace) {
		return p.blockStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) ifStatement() *stmt.If {
	p.consume(token.LeftParen, "Expect '(' after 'if'.")
	condition := p.expression()
	p.consume(token.RightParen, "Expect ')' after if condition.")

	ifStmt := &stmt.If{
		Condition:  condition,
		ThenBranch: p.statement(),
		ElseBranch: nil,
	}
	if p.match(token.Else) {
		ifStmt.ElseBranch = p.statement()
	}

	return ifStmt
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

func (p *Parser) blockStatement() *stmt.Block {
	b := &stmt.Block{
		StatementList: []stmt.Statement{},
	}

	for {
		if p.check(token.RightBrace) || p.isAtEnd() {
			break
		}

		b.StatementList = append(b.StatementList, p.declaration())
	}

	p.consume(token.RightBrace, "Expect '}' after block.")

	return b
}
