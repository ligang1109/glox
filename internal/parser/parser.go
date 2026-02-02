package parser

import (
	"fmt"

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

		statement, err := p.parse()
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

func (p *Parser) parse() (statement stmt.Statement, err error) {
	defer func() {
		if v := recover(); v != nil {
			ok := false
			err, ok = v.(*perror.ParseError)
			if !ok {
				err = fmt.Errorf("Parser.Parse recover from %v", v)
			}
		}
	}()

	return p.declaration(), nil
}

func (p *Parser) init(tokens []*token.Token) {
	p.tokens = tokens
	p.current = 0
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
