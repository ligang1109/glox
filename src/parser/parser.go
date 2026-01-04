package parser

import (
	"glox/expr"
	"glox/perror"
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
	exp := p.comparison()
	for {
		if p.match(token.BangEqual, token.EqualEqual) {
			exp = &expr.Binary{
				Left:     exp,
				Right:    p.comparison(),
				Operator: p.previous(),
			}
		} else {
			break
		}
	}

	return exp
}

// comparison -> term ( ( ">" | ">=" | "<" | "<=" ) term )*
func (p *Parser) comparison() expr.Expr {
	exp := p.term()
	for {
		if p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
			exp = &expr.Binary{
				Left:     exp,
				Right:    p.term(),
				Operator: p.previous(),
			}
		} else {
			break
		}
	}

	return exp
}

// term -> factor ( ( "-" | "+" ) factor )*
func (p *Parser) term() expr.Expr {
	exp := p.factor()
	for {
		if p.match(token.Minus, token.Plus) {
			exp = &expr.Binary{
				Left:     exp,
				Right:    p.factor(),
				Operator: p.previous(),
			}
		} else {
			break
		}
	}

	return exp
}

// factor -> unary ( ( "/" | "*" ) unary )*
func (p *Parser) factor() expr.Expr {
	exp := p.unary()
	for {
		if p.match(token.Slash, token.Star) {
			exp = &expr.Binary{
				Left:     exp,
				Right:    p.unary(),
				Operator: p.previous(),
			}
		} else {
			break
		}
	}

	return exp
}

// unary -> ( "!" | "-" ) unary | primary
func (p *Parser) unary() expr.Expr {
	if p.match(token.Bang, token.Minus) {
		return &expr.Unary{
			Right:    p.unary(),
			Operator: p.previous(),
		}
	}

	return p.primary()
}

// primary -> NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")"
func (p *Parser) primary() expr.Expr {
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
		panic(perror.NewParseError(p.peek(), "Expect ')' after expression."))
	}

	panic(perror.NewParseError(p.peek(), "Unexpected token."))
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
