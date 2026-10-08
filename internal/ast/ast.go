package ast

// Value is a lox value
type Value = any

// VisitResult is what a visitor produces when visiting a node
type VisitResult = any

type Expr interface {
	Accept(visitor ExprVisitor) VisitResult
}

type Stmt interface {
	Accept(visitor StmtVisitor) VisitResult
}

type ExprVisitor interface {
	VisitBinary(binary *Binary) VisitResult
	VisitGrouping(grouping *Grouping) VisitResult
	VisitLiteral(literal *Literal) VisitResult
	VisitUnary(unary *Unary) VisitResult
}

type StmtVisitor interface {
}
