package tree

/*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

/* BINARY TREE PRE-ORDER TRAVERSAL */

func PreOrderTraversal(root *TreeNode) []string {
	var order = make([]string, 0)
	order = traversePreOrder(root, order)
	return order
}

func traversePreOrder(node *TreeNode, order []string) []string {
	if node == nil {
		return order
	}
	order = append(order, node.Val)
	if node.Left != nil {
		order = traversePreOrder(node.Left, order)
	}
	if node.Right != nil {
		order = traversePreOrder(node.Right, order)
	}
	return order
}

/* BINARY TREE IN-ORDER TRAVERSAL */

func InOrderTraversal(root *TreeNode) []string {
	var order = make([]string, 0)
	return traverseInOrder(root, order)
}

func traverseInOrder(node *TreeNode, order []string) []string {
	if node == nil {
		return order
	}
	if node.Left != nil {
		order = traverseInOrder(node.Left, order)
	}
	order = append(order, node.Val)
	if node.Right != nil {
		order = traverseInOrder(node.Right, order)
	}
	return order
}

/* BINARY TREE POST-ORDER TRAVERSAL */

func PostOrderTraversal(root *TreeNode) []string {
	var order = make([]string, 0)
	return traversePostOrder(root, order)
}

func traversePostOrder(node *TreeNode, order []string) []string {
	if node == nil {
		return order
	}
	if node.Left != nil {
		order = traversePostOrder(node.Left, order)
	}
	if node.Right != nil {
		order = traversePostOrder(node.Right, order)
	}
	order = append(order, node.Val)
	return order
}

/* BINARY TREE PRE-ORDER ITERATIVE TRAVERSAL */

func PreOrderTraversalIterative(root *TreeNode) []string {
	var order = make([]string, 0)
	var stack = make([]*TreeNode, 0)
	if root == nil {
		return order
	} else {
		stack = []*TreeNode{root}
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return order
}

// Sum using the generics implementation to create a single sum function
// for both integers and floats.
func Sum[V int | float32](m map[string]V) V {
	var total V
	for _, v := range m {
		total += v
	}
	return total
}

// go is smart enough to determine the generic's type and then
// inform us that x is that type
// var x = Sum(map[string]int{"uno": 1})
