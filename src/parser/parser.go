package parser

import (
	"glox/expr"
	"glox/token"
)

type Parser struct {
	tokens  []*token.Token
	current int
}

func (p *Parser) ParseTokens(tokens []*token.Token) {
	p.init(tokens)
}

func (p *Parser) init(tokens []*token.Token) {
	p.tokens = tokens
	p.current = 0
}

// expression -> equality
func (p *Parser) expression() expr.Expr {
	return p.equality()
}

// equality -> comparison ( ( "!=" | "==" ) comparison )*
func (p *Parser) equality() expr.Expr {
	expr := p.comparison()
	for {

	}
}

// comparison -> term ( ( ">" | ">=" | "<" | "<=" ) term )*
func (p *Parser) comparison() expr.Expr {

}

// term -> factor ( ( "-" | "+" ) factor )*
func (p *Parser) term() expr.Expr {

}

// factor -> unary ( ( "/" | "*" ) unary )*
func (p *Parser) factor() expr.Expr {

}

// unary -> ( "!" | "-" ) unary | primary
func (p *Parser) unary() expr.Expr {

}

// primary -> NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")"
func (p *Parser) primary() expr.Expr {

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
