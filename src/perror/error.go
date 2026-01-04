package perror

import (
	"fmt"
	"glox/token"
)

type ParseError struct {
	token   *token.Token
	message string
}

func NewParseError(token *token.Token, message string) *ParseError {
	return &ParseError{
		token:   token,
		message: message,
	}
}

func (e *ParseError) Error() string {
	var pos string
	if e.token.Type == token.Eof {
		pos = "end"
	} else {
		pos = fmt.Sprintf("'%s'", e.token.Lexeme)
	}

	return fmt.Sprintf("line %d at %s, %s", e.token.Line, pos, e.message)
}
