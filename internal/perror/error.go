package perror

import (
	"fmt"

	"github.com/ligang1109/glox/pkg/token"
)

type tokenError struct {
	token   *token.Token
	prefix  string
	message string
}

func newTokenError(token *token.Token, prefix, message string) *tokenError {
	return &tokenError{
		token:   token,
		message: message,
		prefix:  prefix,
	}
}

func (e *tokenError) Error() string {
	if e.token == nil {
		return fmt.Sprintf("%s, %s", e.prefix, e.message)
	}

	var pos string
	if e.token.Type == token.Eof {
		pos = "end"
	} else {
		pos = fmt.Sprintf("'%s'", e.token.Lexeme)
	}

	return fmt.Sprintf("%s, line %d at %s, %s", e.prefix, e.token.Line, pos, e.message)
}

type ParseError struct {
	*tokenError
}

func NewParseError(token *token.Token, message string) *ParseError {
	return &ParseError{
		tokenError: newTokenError(token, "ParseError", message),
	}
}

type RuntimeError struct {
	*tokenError
}

func NewRuntimeError(token *token.Token, message string) *RuntimeError {
	return &RuntimeError{
		tokenError: newTokenError(token, "RuntimeError", message),
	}
}
