package expr

type Literal struct {
	Value any
}

func (l *Literal) Name() string {
	return "Literal"
}
