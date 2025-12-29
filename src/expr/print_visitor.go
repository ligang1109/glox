package expr

import (
	"fmt"
)

type PrintVisitor struct {
	curRow    int
	rowColMap map[int]int
}

func NewPrintVisitor() *PrintVisitor {
	return &PrintVisitor{
		rowColMap: map[int]int{},
	}
}

func (v *PrintVisitor) assignNodeCol(row int) int {
	col, ok := v.rowColMap[row]
	if !ok {
		col = 0
	} else {
		col++
	}

	v.rowColMap[row] = col

	return col
}

func (v *PrintVisitor) visitExpr(exp Expr) {
}

func (v *PrintVisitor) VisitBinary(binary *Binary) {
	name := binary.Name()
	graph := fmt.Sprintf("%s --> %s")
}

func (v *PrintVisitor) VisitGrouping(grouping *Grouping) {

}

func (v *PrintVisitor) VisitLiteral(literal *Literal) {

}

func (v *PrintVisitor) VisitUnary(unary *Unary) {

}
