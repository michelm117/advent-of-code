package day_1

import (
	"fmt"
	"strings"

	"github.com/michelm117/advent-of-code/utils"
)

type Tree struct {
	Tokens *Node
}

func (t *Tree) AddPattern(pattern string) {
	if t.Tokens == nil {
		t.Tokens = &Node{Value: ""}
	}

	t.addPatternRecursive(t.Tokens, pattern)
}

func (t *Tree) addPatternRecursive(node *Node, pattern string) {
	if len(pattern) == 0 {
		return
	}

	for _, child := range node.Nodes {
		if string(pattern[0]) == child.Value {
			t.addPatternRecursive(child, pattern[len(child.Value):])
			return
		}
	}
	newNode := node.AddChild(string(pattern[0]))
	t.addPatternRecursive(newNode, pattern[1:])
}

func (t *Tree) Print() {
	s := t.printRecursive(t.Tokens, "")
	fmt.Println(s)
}

func (t *Tree) printRecursive(node *Node, indent string) string {
	s := "" + indent + node.Value + "\n"
	for _, child := range node.Nodes {
		s += t.printRecursive(child, indent+"  ")
	}

	return s
}

func (t *Tree) GetAllPatterns() []string {
	return t.getAllPatternsRecursive(t.Tokens, "")
}

func (t *Tree) getAllPatternsRecursive(node *Node, prefix string) []string {
	if node == nil {
		return []string{}
	}

	if len(node.Nodes) == 0 {
		return []string{prefix + node.Value}
	}

	var patterns []string
	for _, child := range node.Nodes {
		patterns = append(patterns, t.getAllPatternsRecursive(child, prefix+node.Value)...)
	}
	return patterns
}

func (t *Tree) IsPrefix(prefix string) bool {
	return t.isPrefixRecursive(t.Tokens, strings.ToUpper(prefix)) != nil
}

func (t *Tree) isPrefixRecursive(node *Node, prefix string) *Node {
	if len(prefix) == 0 {
		return node
	}

	for _, child := range node.Nodes {
		if strings.HasPrefix(prefix, child.Value) {
			return t.isPrefixRecursive(child, prefix[len(child.Value):])
		}
	}
	return nil
}

func (t *Tree) Reverse() *Tree {
	reversedTree := Tree{}
	for _, pattern := range t.GetAllPatterns() {
		reversedString := utils.ReverseString(pattern)
		reversedTree.AddPattern(reversedString)
	}
	return &reversedTree
}
