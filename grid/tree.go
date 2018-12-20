package grid

import (
	"fmt"
)

type Node struct{
	index int
	parent *Node
	children []*Node
	Value Cell
}

func NewNode(c Cell, p *Node) *Node {
	n := &Node{
		parent: p,
		children: []*Node{},
		Value: c,
	}
	p.children = append(p.children, n)
	return n
}

func (n *Node) Path() []Cell {
	path := []Cell{n.Value}

	p := n.parent
	for p != nil {
		path = append([]Cell{p.Value}, path...)
		p = p.parent
	}

	return path
}

func (n *Node) Pos() Coordinate {
	return n.Value.Pos()
}

func (n *Node) String() string {
	p := "root"
	if n.parent != nil {
		p = "node"
	}
	return fmt.Sprintf("<P[%s] {%d,%d}>", p, n.Value.Pos().X, n.Value.Pos().Y)
}

type Tree []*Node

func NewTree(root Cell) *Tree {
	t := make(Tree, 1)
	t[0] = &Node{
		parent: nil,
		children: []*Node{},
		Value: root,
	}
	return &t
}

func (t *Tree) Push(n *Node) {
	n.index = len(*t)
	*t = append(*t, n)
}

func (t *Tree) Pop() *Node {
	old := *t
	l := len(old)
	n := old[l-1]
	n.index = -1
	*t = old[0:l-1]
	return n
}

type TreeByReadOrder Tree

func (t TreeByReadOrder) Len() int {
	return len(t)
}

func (t TreeByReadOrder) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
	t[i].index = i
	t[j].index = j
}

func (t TreeByReadOrder) Less(i, j int) bool {
	if t[i].Value.Pos().Y == t[j].Value.Pos().Y {
		return t[i].Value.Pos().X < t[j].Value.Pos().X
	}
	return t[i].Value.Pos().Y < t[j].Value.Pos().Y
}

func (t *TreeByReadOrder) Push(x interface{}) {
	(*Tree)(t).Push(x.(*Node))
}

func (t *TreeByReadOrder) Pop() interface{} {
	return (*Tree)(t).Pop()
}