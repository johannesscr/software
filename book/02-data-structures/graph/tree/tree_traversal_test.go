package tree

import (
	"fmt"
	"testing"
)

var root = TreeNode{
	Val: "F",
	Left: &TreeNode{
		Val: "B",
		Left: &TreeNode{
			Val: "A",
		},
		Right: &TreeNode{
			Val: "D",
			Left: &TreeNode{
				Val: "C",
			},
			Right: &TreeNode{
				Val: "E",
			},
		},
	},
	Right: &TreeNode{
		Val: "G",
		Right: &TreeNode{
			Val: "I",
			Left: &TreeNode{
				Val: "H",
			},
		},
	},
}

func TestPreOrderTraversal(t *testing.T) {
	o := PreOrderTraversal(&root)
	output := []string{"F", "B", "A", "D", "C", "E", "G", "I", "H"}
	output1 := fmt.Sprintf("%v", o)
	output2 := fmt.Sprintf("%v", output)
	if output1 != output2 {
		t.Errorf("expected output %s got %s", output2, output1)
	}
}

func TestInOrderTraversal(t *testing.T) {
	o := InOrderTraversal(&root)
	output := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	output1 := fmt.Sprintf("%v", o)
	output2 := fmt.Sprintf("%v", output)
	if output1 != output2 {
		t.Errorf("expected output %s got %s", output2, output1)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	o := PostOrderTraversal(&root)
	output := []string{"A", "C", "E", "D", "B", "H", "I", "G", "F"}
	output1 := fmt.Sprintf("%v", o)
	output2 := fmt.Sprintf("%v", output)
	if output1 != output2 {
		t.Errorf("expected output %s got %s", output2, output1)
	}
}

func TestPreOrderTraversalIterative(t *testing.T) {
	o := PreOrderTraversalIterative(&root)
	output := []string{"F", "B", "A", "D", "C", "E", "G", "I", "H"}
	output1 := fmt.Sprintf("%v", o)
	output2 := fmt.Sprintf("%v", output)
	if output1 != output2 {
		t.Errorf("expected output %s got %s", output2, output1)
	}
}
