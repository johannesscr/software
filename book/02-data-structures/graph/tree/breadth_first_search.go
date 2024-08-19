package tree

import (
	"fmt"
	"strings"
)

type Node struct {
	Val      string
	Children []*Node
}

func (n *Node) Print() {
	children := make([]string, 0)
	for _, child := range n.Children {
		children = append(children, child.Val)
	}
	fmt.Printf("%s -> %s\n", n.Val, strings.Join(children, " - "))
}

func PrintQueue(xn []*Node) {
	queueVals := make([]string, 0)
	for _, node := range xn {
		queueVals = append(queueVals, node.Val)
	}
	fmt.Printf("q: %s\n", strings.Join(queueVals, ","))
}

// BFS is breadth first search and starts from the given node (as the root
// node) to start doing a breadth first search
func (n *Node) BFS(searchValue string) []string {
	q := []*Node{n} // q denotes queue
	values := make([]string, 0)

	for len(q) > 0 {
		PrintQueue(q)
		// dequeue the queue
		node := q[0]
		node.Print()
		q = q[1:]
		values = append(values, node.Val)
		if searchValue == node.Val {
			return values
		}
		q = append(q, node.Children...)
	}
	return []string{}
}

func (n *Node) DFS(searchValue string) []string {
	var nodes []string
	nodes = traverse(nodes, n, searchValue)
	return nodes
}

func traverse(nodes []string, node *Node, value string) []string {
	node.Print()
	nodes = append(nodes, node.Val)
	if node.Val == value {
		return nodes
	}
	for _, n := range node.Children {
		xn := traverse(nodes, n, value)
		if xn != nil {
			return xn
		}
	}
	return nil
}
