package challenges

import (
	"fmt"
	"testing"
)

/*
You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and
'-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and
concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates
to target.



Example 1:

Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
Example 2:

Input: nums = [1], target = 1
Output: 1


Constraints:

- `1 <= nums.length <= 20`
- `0 <= nums[i] <= 1000`
- `0 <= sum(nums[i]) <= 1000`
- `-1000 <= target <= 1000`
*/

const (
	ADD = iota
	SUBTRACT
)

/* SOLUTION 1 */

func findTargetSumWays2(nums []int, target int) int {
	permutations := make(map[string]int, 0)

	dfsTargetSum("", 0, nums, target, &permutations)
	return len(permutations)
}

func nextPermutationLevel(val, total int, variant int) (string, int) {
	switch variant {
	case ADD:
		total = total + val
		return fmt.Sprintf("+%d", val), total
	case SUBTRACT:
		total = total - val
		return fmt.Sprintf("-%d", val), total
	default:
		return "?", total
	}
}

func dfsTargetSum(permutation string, total int, nums []int, target int, permutations *map[string]int) {
	// termination condition
	if len(nums) == 0 {
		// we are done with the depth first search
		if total == target {
			// we need to add to the value not the address
			(*permutations)[permutation] = total
		}
		return
	}

	val := nums[0]
	nums = nums[1:]
	// add value
	addPerm, totalAdd := nextPermutationLevel(val, total, ADD)
	addPermutation := fmt.Sprintf("%s%s", permutation, addPerm)
	// next depth
	dfsTargetSum(addPermutation, totalAdd, nums, target, permutations)

	// subtract value
	subPerm, totalSub := nextPermutationLevel(val, total, SUBTRACT)
	subPermutation := fmt.Sprintf("%s%s", permutation, subPerm)
	// next depth
	dfsTargetSum(subPermutation, totalSub, nums, target, permutations)
}

/* SOLUTION 2 */

type SolutionMap map[string]int
type IndexMap map[int]SolutionMap

func findTargetSumWays(nums []int, target int) int {
	var indexMap IndexMap
	for index, _ := range nums {
		indexMap[index] = SolutionMap{}
		// remove the old memory
		if _, ok := indexMap[index-2]; ok {
			delete(indexMap, index-2)
		}
	}
	return 0
}

//func TestFindTargetSumWays(t *testing.T) {
//	tests := []struct {
//		nums   []int
//		target int
//		out    int
//	}{
//		{
//			nums:   []int{1, 1, 1, 1, 1},
//			target: 3,
//			out:    5,
//		},
//		{
//			nums:   []int{27, 22, 39, 22, 40, 32, 44, 45, 46, 8, 8, 21, 27, 8, 11, 29, 16, 15, 41, 0},
//			target: 10,
//			out:    10,
//		},
//		{
//			nums:   []int{2, 20, 24, 38, 44, 21, 45, 48, 30, 48, 14, 9, 21, 10, 46, 46, 12, 48, 12, 38},
//			target: 48,
//			out:    10,
//		},
//	}
//
//	for _, tc := range tests {
//		t.Run("", func(t *testing.T) {
//			o := findTargetSumWays(tc.nums, tc.target)
//			if o != tc.out {
//				t.Errorf("expected number of ways %d got %d", tc.out, o)
//			}
//		})
//	}
//}

/* SOLUTION 3 */

/*
Use a tree to step through all the solutions.
*/

type tree struct {
	nodes []*node
}

func newTree() tree {
	root := &node{
		index: 0,
		val:   0,
	}
	return tree{
		nodes: []*node{root},
	}
}

func (t *tree) add(index, num int) {
	// add the initial node
	for _, leaf := range t.leaves() {
		leftNode := &node{
			index:    index,
			val:      leaf.val - num,
			solution: fmt.Sprintf("%s-%d", leaf.solution, num),
		}
		rightNode := &node{
			index:    index,
			val:      leaf.val + num,
			solution: fmt.Sprintf("%s+%d", leaf.solution, num),
		}
		leaf.left = leftNode
		leaf.right = rightNode
		t.nodes = append(t.nodes, leftNode, rightNode)
	}
}

func (t *tree) leaves() []*node {
	leaves := make([]*node, 0)
	for _, n := range t.nodes {
		if n.isLeaf() {
			leaves = append(leaves, n)
		}
	}
	t.nodes = leaves
	return t.nodes
}

type node struct {
	index    int
	val      int
	solution string
	left     *node
	right    *node
}

// IsLeaf returns true if the node is a leaf node. 
func (n node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func findTargetSumWays3(nums []int, target int) int {
	solTree := newTree()

	for index, n := range nums {
		solTree.add(index, n)
	}

	sols := solTree.leaves()

	solutions := make([]*node, 0)

	for _, s := range sols {
		if s.val == target {
			solutions = append(solutions, s)
		}
	}

	return len(solutions)
}

func TestFindTargetSumWays3(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		out    int
	}{
		{
			nums:   []int{1},
			target: 1,
			out:    1,
		},
		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			out:    5,
		},
		{
			nums:   []int{10, 9, 6, 4, 19, 0, 41, 30, 27, 15, 14, 39, 33, 7, 34, 17, 24, 46, 2, 46},
			target: 45,
			out:    6606,
		},
		//		{
		//			nums:   []int{27, 22, 39, 22, 40, 32, 44, 45, 46, 8, 8, 21, 27, 8, 11, 29, 16, 15, 41, 0},
		//			target: 10,
		//			out:    10,
		//		},
		//		{
		//			nums:   []int{2, 20, 24, 38, 44, 21, 45, 48, 30, 48, 14, 9, 21, 10, 46, 46, 12, 48, 12, 38},
		//			target: 48,
		//			out:    10,
		//		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			o := findTargetSumWays3(tc.nums, tc.target)
			if o != tc.out {
				t.Errorf("expected number of ways %d got %d", tc.out, o)
			}
		})
	}
}
