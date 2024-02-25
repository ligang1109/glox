package scanner

import (
	"fmt"

	"glox/log"
	"glox/token"
)

type Scanner struct {
	source string
	tokens []*token.Token

	start   int
	current int
	line    int

	hasError bool
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

func (s *Scanner) HasError() bool {
	return s.hasError
}

func (s *Scanner) init(source string) {
	s.source = source
	s.tokens = []*token.Token{}

	s.start = 0
	s.current = 0
	s.line = 1

	s.hasError = false
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) error(msg string) {
	s.hasError = true
	log.Logger.Error(fmt.Sprintf("line %d error: %s", s.line, msg))
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
	case "!":
		if s.match("=") {
			s.addToken(token.BangEqual, nil)
		} else {
			s.addToken(token.Bang, nil)
		}
	case "=":
		if s.match("=") {
			s.addToken(token.EqualEqual, nil)
		} else {
			s.addToken(token.Equal, nil)
		}
	case "<":
		if s.match("=") {
			s.addToken(token.LessEqual, nil)
		} else {
			s.addToken(token.Less, nil)
		}
	case ">":
		if s.match("=") {
			s.addToken(token.GreaterEqual, nil)
		} else {
			s.addToken(token.Greater, nil)
		}
	case "/":
		if s.match("/") {
			for {
				if s.peek() != "\n" && !s.isAtEnd() {
					s.advance()
				} else {
					break
				}
			}
		} else {
			s.addToken(token.Slash, nil)
		}
	case " ", "\r", "\t":
	case "\n":
		s.line++
	case `"`:
		s.string()
	default:
		s.error("Unexpected character.")
	}
}

func (s *Scanner) advance() string {
	c := s.charAt(s.current)
	s.current++

	return c
}

func (s *Scanner) charAt(pos int) string {
	return string(s.source[pos])
}

func (s *Scanner) addToken(tokenType token.Type, literal any) {
	s.tokens = append(s.tokens, &token.Token{
		TokenType: tokenType,
		Lexeme:    s.source[s.start:s.current],
		Literal:   literal,
		Line:      s.line,
	})
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}

	if s.charAt(s.current) != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return ""
	}

	return s.charAt(s.current)
}

func (s *Scanner) string() {
	for {
		c := s.peek()
		if c != `"` && !s.isAtEnd() {
			if c == "\n" {
				s.line++
			}
			s.advance()
		} else {
			break
		}
	}

	if s.isAtEnd() {
		s.error("Unterminated string.")
		return
	}

	// The closing "
	s.advance()

	s.addToken(token.String, s.source[s.start+1:s.current-1])
}
