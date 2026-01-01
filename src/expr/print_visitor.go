package expr

import (
	"fmt"
	"strings"
)

type printVisitorNode struct {
	row  int
	col  int
	text string

	parent *printVisitorNode
}

func (n *printVisitorNode) nodeId() string {
	return fmt.Sprintf("r%dc%d", n.row, n.col)
}

func (n *printVisitorNode) drawNode() string {
	return fmt.Sprintf("%s[%s]", n.nodeId(), n.text)
}

func (n *printVisitorNode) drawConnect() string {
	return fmt.Sprintf("%s --> %s", n.parent.nodeId(), n.nodeId())
}

type PrintVisitor struct {
	lineList []string

	curNode      *printVisitorNode
	colAssignMap map[int]int
}

func NewPrintVisitor() *PrintVisitor {
	return &PrintVisitor{
		colAssignMap: map[int]int{},
	}
}

func (v *PrintVisitor) assignCol(row int) int {
	col, ok := v.colAssignMap[row]
	if !ok {
		col = 0
	} else {
		col++
	}

	v.colAssignMap[row] = col

	return col
}

func (v *PrintVisitor) newNode(text string, parent *printVisitorNode) *printVisitorNode {
	row := 0
	if parent != nil {
		row = parent.row + 1
	}

	return &printVisitorNode{
		row:  row,
		col:  v.assignCol(row),
		text: text,

		parent: parent,
	}
}

func (v *PrintVisitor) drawNode(node *printVisitorNode) {
	v.lineList = append(v.lineList, node.drawNode())
	if node.parent != nil {
		v.lineList = append(v.lineList, node.drawConnect())
	}
}

func (v *PrintVisitor) drawChild(child *printVisitorNode, exp Expr) {
	v.curNode = child
	exp.Accept(v)
	v.curNode = child.parent
}

func (v *PrintVisitor) Graph() string {
	graph := "flowchart TD\n"
	graph += strings.Join(v.lineList, "\n")
	graph += "\n"

	return graph
}

func (v *PrintVisitor) VisitBinary(binary *Binary) {
	node := v.curNode
	if node == nil {
		node = v.newNode(binary.Name(), nil)
	}
	v.drawNode(node)

	child := v.newNode(binary.Left.Name(), node)
	v.drawChild(child, binary.Left)

	child = v.newNode(fmt.Sprintf("Operator %s", binary.Operator.Lexeme), node)
	v.drawNode(child)

	child = v.newNode(binary.Right.Name(), node)
	v.drawChild(child, binary.Right)
}

func (v *PrintVisitor) VisitGrouping(grouping *Grouping) {
	text := grouping.Expression.Name()
	child := v.curNode
	if child == nil {
		child = v.newNode(text, nil)
	} else {
		child.text = text
	}

	v.drawChild(child, grouping.Expression)
}

func (v *PrintVisitor) VisitLiteral(literal *Literal) {
	text := fmt.Sprintf("%s %v", literal.Name(), literal.Value)
	node := v.curNode
	if node == nil {
		node = v.newNode(text, nil)
	} else {
		node.text = text
	}

	v.drawNode(node)
}

func (v *PrintVisitor) VisitUnary(unary *Unary) {
	node := v.curNode
	if node == nil {
		node = v.newNode(unary.Name(), nil)
	}
	v.drawNode(node)

	child := v.newNode(fmt.Sprintf("Operator %s", unary.Operator.Lexeme), node)
	v.drawNode(child)

	child = v.newNode(unary.Right.Name(), node)
	v.drawChild(child, unary.Right)
}
