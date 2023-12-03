package day_1

import "strings"

type Node struct {
	Value string
	Nodes []*Node
}

func (n *Node) AddChild(pattern string) *Node {
	if n.Nodes == nil {
		n.Nodes = []*Node{}
	}

	newNode := &Node{Value: strings.ToUpper(pattern)}
	n.Nodes = append(n.Nodes, newNode)
	return newNode
}
