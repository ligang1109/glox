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

func (v *PrintVisitor) drawChild(child *printVisitorNode, exp Expression) {
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

func (v *PrintVisitor) VisitBinaryExpr(binary *Binary) {
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

func (v *PrintVisitor) VisitGroupingExpr(grouping *Grouping) {
	text := grouping.Expression.Name()
	child := v.curNode
	if child == nil {
		child = v.newNode(text, nil)
	} else {
		child.text = text
	}

	v.drawChild(child, grouping.Expression)
}

func (v *PrintVisitor) VisitLiteralExpr(literal *Literal) {
	text := fmt.Sprintf("%s %v", literal.Name(), literal.Value)
	node := v.curNode
	if node == nil {
		node = v.newNode(text, nil)
	} else {
		node.text = text
	}

	v.drawNode(node)
}

func (v *PrintVisitor) VisitUnaryExpr(unary *Unary) {
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

func (v *PrintVisitor) VisitVariableExpr(variable *Variable) {
	text := fmt.Sprintf("%s %s", variable.Name(), variable.VarName.Lexeme)
	node := v.curNode
	if node == nil {
		node = v.newNode(text, nil)
	} else {
		node.text = text
	}

	v.drawNode(node)
}

func (v *PrintVisitor) VisitAssignExpr(assign *Assign) {
	node := v.curNode
	if node == nil {
		node = v.newNode(assign.Name(), nil)
	}
	v.drawNode(node)

	child := v.newNode(fmt.Sprintf("var %s =", assign.VarName.Lexeme), node)
	v.drawNode(child)

	child = v.newNode(assign.Value.Name(), node)
	v.drawChild(child, assign.Value)
}

func (v *PrintVisitor) VisitLogicalExpr(logical *Logical) {
	node := v.curNode
	if node == nil {
		node = v.newNode(logical.Name(), nil)
	}
	v.drawNode(node)

	child := v.newNode(logical.Left.Name(), node)
	v.drawChild(child, logical.Left)

	child = v.newNode(fmt.Sprintf("Operator %s", logical.Operator.Lexeme), node)
	v.drawNode(child)

	child = v.newNode(logical.Right.Name(), node)
	v.drawChild(child, logical.Right)
}

func (v *PrintVisitor) VisitCallExpr(call *Call) {
	callNode := v.curNode
	if callNode == nil {
		callNode = v.newNode(call.Name(), nil)
	}
	v.drawNode(callNode)

	calleeNode := v.newNode("Callee", callNode)
	v.drawNode(calleeNode)

	child := v.newNode(call.Callee.Name(), calleeNode)
	v.drawChild(child, call.Callee)
	v.curNode = callNode

	argumentsNode := v.newNode("Arguments", callNode)
	v.drawNode(argumentsNode)
	for _, arg := range call.Arguments {
		child := v.newNode(arg.Name(), argumentsNode)
		v.drawChild(child, arg)
	}
	v.curNode = callNode
}
