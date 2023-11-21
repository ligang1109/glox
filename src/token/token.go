package token

const (
	// Single-character tokens.
	TypeLeftParen  = "LEFT_PAREN"
	TypeRightParen = "RIGHT_PAREN"
	TypeLeftBrace  = "LEFT_BRACE"
	TypeRightBrace = "RIGHT_BRACE"
	TypeComma      = "COMMA"
	TypeDot        = "DOT"
	TypeMinus      = "MINUS"
	TypePlus       = "PLUS"
	TypeSemicolon  = "SEMICOLON"
	TypeSlash      = "SLASH"
	TypeStar       = "STAR"

	// One or two character tokens.
	TypeBang         = "BANG"
	TypeBangEqual    = "BANG_EQUAL"
	TypeEqual        = "EQUAL"
	TypeEqualEqual   = "EQUAL_EQUAL"
	TypeGreater      = "GREATER"
	TypeGreaterEqual = "GREATER_EQUAL"
	TypeLess         = "LESS"
	TypeLessEqual    = "LESS_EQUAL"
)
