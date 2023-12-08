package scanner

import (
	"glox/token"
)

type Scanner struct {
	source string
	tokens []*token.Token

	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens(source string) []*token.Token {
	s.init(source)

	for {
		if s.isAtEnd() {
			break
		}

		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, &token.Token{
		TokenType: token.Eof,
		Lexeme:    "",
		Literal:   nil,
		Line:      s.line,
	})

	return s.tokens
}

func (s *Scanner) init(source string) {
	s.source = source
	s.tokens = []*token.Token{}

	s.start = 0
	s.current = 0
	s.line = 1
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(token.LeftParen, nil)
	case ")":
		s.addToken(token.RightParen, nil)
	case "{":
		s.addToken(token.LeftBrace, nil)
	case "}":
		s.addToken(token.RightBrace, nil)
	case ",":
		s.addToken(token.Comma, nil)
	case ".":
		s.addToken(token.Dot, nil)
	case "-":
		s.addToken(token.Minus, nil)
	case "+":
		s.addToken(token.Plus, nil)
	case ";":
		s.addToken(token.Semicolon, nil)
	case "*":
		s.addToken(token.Star, nil)
	}
}

func (s *Scanner) advance() string {
	c := string(s.source[s.current])
	s.current++

	return c
}

func (s *Scanner) addToken(tokenType token.Type, literal any) {
	s.tokens = append(s.tokens, &token.Token{
		TokenType: tokenType,
		Lexeme:    s.source[s.start : s.current+1],
		Literal:   literal,
		Line:      s.line,
	})
}
