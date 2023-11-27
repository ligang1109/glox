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

}

func (s *Scanner) advance() string {

}
