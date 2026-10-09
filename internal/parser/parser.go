package parser

import (
	"github.com/ligang1109/glox/internal/perror"
	"github.com/ligang1109/glox/internal/token"
)

type Parser struct {
	tokens    []*token.Token
	current   int
	errorList []error
}

func (p *Parser) init(tokens []*token.Token) {
	p.tokens = tokens
	p.current = 0
	p.errorList = []error{}
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
	return p.peek().Type == tokenType
}

func (p *Parser) isAtEnd() bool {
	return p.check(token.Eof)
}

func (p *Parser) peek() *token.Token {
	return p.tokens[p.current]
}

func (p *Parser) advance() *token.Token {
	tok := p.peek()
	if !p.isAtEnd() {
		p.current++
	}

	return tok
}

func (p *Parser) previous() *token.Token {
	return p.tokens[(p.current - 1)]
}

func (p *Parser) consume(tokenType token.Type, message string) (*token.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}

	return nil, p.error(nil, message)
}

func (p *Parser) error(tok *token.Token, message string) error {
	if tok == nil {
		tok = p.peek()
	}

	return perror.NewParseError(tok, message)
}
