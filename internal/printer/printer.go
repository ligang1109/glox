package printer

import (
	"fmt"
	"strings"

	"github.com/ligang1109/glox/internal/ast"
)

type Printer struct {
	count    int
	lineList []string
}

type node struct {
	label    string
	children []*node
}

func (p *Printer) reset() {
	p.count = 0
	p.lineList = []string{}
}

func (p *Printer) exprToNode(expr ast.Expr) *node {
	n, _ := expr.Accept(p).(*node)

	return n
}

func (p *Printer) draw(n *node, parentID string) {
	id := fmt.Sprintf("n%d", p.count)
	p.count++

	p.lineList = append(p.lineList, fmt.Sprintf(`%s["%s"]`, id, p.escapeLabel(n.label)))
	if parentID != "" {
		p.lineList = append(p.lineList, fmt.Sprintf("%s --> %s", parentID, id))
	}
	for _, child := range n.children {
		p.draw(child, id)
	}
}

// escapeLabel makes the raw label safe inside a quoted mermaid string.
func (p *Printer) escapeLabel(s string) string {
	return strings.ReplaceAll(s, `"`, "#quot;")
}

func (p *Printer) render(n *node) string {
	p.draw(n, "")

	graph := "flowchart TD\n"
	graph += strings.Join(p.lineList, "\n")
	graph += "\n"

	return graph
}

func (p *Printer) PrintExpr(expr ast.Expr) string {
	p.reset()

	return p.render(p.exprToNode(expr))
}

func (p *Printer) PrintStmt(stmt ast.Stmt) string {
	return ""
}
