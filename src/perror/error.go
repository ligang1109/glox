package perror

import "glox/token"

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
