package ast

type Node interface {
	Name() string
}

type Value any

type Expr interface {
	Node

	Accept(visitor ExprVisitor) Value
}

type Stmt interface {
	Node

	Accept(visitor StmtVisitor)
}

type ExprVisitor interface {
	VisitBinary(binary *Binary) Value
	VisitGrouping(grouping *Grouping) Value
	VisitLiteral(literal *Literal) Value
	VisitUnary(unary *Unary) Value
}

type StmtVisitor interface {
}
