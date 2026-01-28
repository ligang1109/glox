package interpreter

import (
	"fmt"

	"github.com/ligang1109/glox/internal/perror"
	"github.com/ligang1109/glox/pkg/token"
)

type Environment struct {
	enclosing *Environment
	valueMap  map[string]any
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		enclosing: enclosing,
		valueMap:  map[string]any{},
	}
}

func (e *Environment) Define(name string, value any) {
	e.valueMap[name] = value
}

func (e *Environment) Value(name *token.Token) any {
	v, ok := e.valueMap[name.Lexeme]
	if ok {
		return v
	}

	if e.enclosing == nil {
		e.error(fmt.Sprintf("Undefined variable %s.", name.Lexeme))
	}

	return e.enclosing.Value(name)
}

func (e *Environment) Assign(name *token.Token, value any) {
	_, ok := e.valueMap[name.Lexeme]
	if ok {
		e.valueMap[name.Lexeme] = value
		return
	}

	if e.enclosing == nil {
		e.error(fmt.Sprintf("Undefined variable %s.", name.Lexeme))
	}

	e.enclosing.Assign(name, value)
}

func (e *Environment) error(message string) {
	panic(perror.NewRuntimeError(nil, message))
}
