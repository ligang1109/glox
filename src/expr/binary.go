package expr

import "glox/token"

type Binary struct {
	Left     Expr
	Right    Expr
	Operator *token.Token
}

func (b *Binary) Name() string {
	return "Binary"
}
