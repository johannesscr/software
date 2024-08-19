package challenges

import (
	"testing"
)

/*
You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots:
'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely
and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each
move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4
wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of
these codes, the wheels of the lock will stop turning and you will be unable to
open it.

Given a target representing the value of the wheels that will unlock the lock,
return the minimum total number of turns required to open the lock, or -1 if it
is impossible.



Example 1:

Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" ->
"1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would
be invalid, because the wheels of the lock become stuck after the display
becomes the dead end "0102".

Example 2:

Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".

Example 3:

Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
	target = "8888"
Output: -1
Explanation: We cannot reach the target without getting stuck.


Constraints:

1 <= deadends.length <= 500
deadends[i].length == 4
target.length == 4
target will not be in the list deadends.
target and deadends[i] consist of digits only
*/

/*
**Intuition**

We can think of this problem as a shortest path problem on a graph: there are
`10000` nodes (strings `'0000'` to `'9999'`), and there is and edge between two
nodes if they differ in one digit, that digit differs by 1 (wrapping around, so
`'0'` and `'9'` differ by 1), and if both nodes are not in `deadends`.

**Algorithm**

To solve a shortest path problem, we use a breadth-first-search. The basic
structure uses a Queue `queue` plus a Set `seen` that records if a node has ever
been enqueued. This pushed all the work to the `neighbours` function - in our
Python implementation, all the code after `while enqueue:` is template code
that, except for `if node in dead: continue`.

As for the `neighbours` function, for each position in the lock `i = 0, 1, 2, 3`
and for each of the turns `d = -1, 1`, we determine the value fo the lock after
the `i`-th wheel has been turned in the direction `d`.

Care should be taken in our algorithm, as the graph does not have an edge unless
both nodes are not in `deadends`. If our `neighbours` function checks only the
`target` for being in `deadends`, we also need ot check whether `'0000'` is in
`deadends` at the beginning. In our implementation, we check at the beginning.
In out implementation, we check at the visitor level to neatly handle this
problem in all cases.

**Complexity Analysis**

- Time complexity: `O(N^2 * A^N + D)`, where `A` is the number of digits in our
  alphabet, `N` is the number of digits in the lock, and `D` is the size of
  `deadends`. We might visit every lock combination, plus we need to instantiate
  our `deadends` Set. When we visit every lock combination, we spend `O(N^2)`
  time through and constructing each node.
- Space complexity: `O(A^N + D)`, for the queue and the set `seen`.
*/

const (
	INCREASE = iota
	DECREASE
)

type Step struct {
	node  string
	level int
}

// openLock returns the minimum total number of turns required to open the lock,
// or -1 if it is impossible. The lock initially starts at '0000'.
func openLock(deadends []string, target string) int {
	combinations := []Step{
		{"0000", 0},
	}
	seen := make(map[string]bool, 0)
	deadEnds := make(map[string]bool, 0)

	for _, deadEnd := range deadends {
		deadEnds[deadEnd] = true
	}

	for len(combinations) > 0 {
		step := combinations[0]
		// check if the step opens the lock
		if step.node == target {
			return step.level
		}
		// only update the queue now that we have not yet found the target
		combinations = combinations[1:]
		// check if the current node is in the set of deadends
		if isDead := deadEnds[step.node]; isDead {
			continue
		}

		for _, neighbour := range neighbours(step.node, step.level+1) {
			if isSeen := seen[neighbour.node]; !isSeen {
				// add the new neighbour to
				combinations = append(combinations, neighbour)
				// update the hash table of seen nodes
				seen[neighbour.node] = true
			}
		}
	}
	return -1
}

// neighbours returns the neighbours of the node.
func neighbours(node string, level int) []Step {
	// convert the string into an array of 4 integers
	// then we can turn the "wheel" either up (+1) or down (-1).
	// note the wheel wraps around from 9->1 and from 1->9.

	// the next turn denotes the current node, where every combination is tried,
	// that is, each wheel is turned once in a positive or negative direction.
	var combinations = make([]Step, 0)

	for i, val := range node {
		// val is now a rune
		valIncrease := turnWheel(val, INCREASE)
		valDecrease := turnWheel(val, DECREASE)
		stepUp := Step{
			node:  replaceAtIndex(node, valIncrease, i),
			level: level,
		}
		stepDown := Step{
			node:  replaceAtIndex(node, valDecrease, i),
			level: level,
		}

		combinations = append(combinations, stepUp, stepDown)
	}

	return combinations
}

// replaceAtIndex takes the node string such as "0000" and replaces the
// character at index `i` with the rune `r` value.
//
// Example:
// replaceAtIndex("0102", 51, 2)
// >>> "0132"
func replaceAtIndex(node string, r rune, i int) string {
	newNode := []rune(node)
	newNode[i] = r
	return string(newNode)
}

// turnWheel takes a rune value for the current integer value on the wheel dial
// then simulates an incremental turn in both directions to increase and
// decrease the wheel count.
//
// Note: the wheel rotates in a circular pattern such that
//
// ... 8 -> 9 -> 0 -> 1 -> 2 ->  ... -> 2 -> 1 -> 0 -> 9 -> 8 ...
func turnWheel(value rune, inc int) rune {
	if inc == INCREASE {
		if value == 57 { // "9"
			return 48 // "0"
		}
		return value + 1
	}
	if inc == DECREASE {
		if value == 48 { // "0"
			return 57 // "9"
		}
		return value - 1
	}
	return 48
}

func TestNeighbours(t *testing.T) {
	tests := []struct {
		in    string
		level int
		out   []string
	}{
		{
			in:    "0000",
			level: 2,
			out: []string{
				"1000",
				"0100",
				"0010",
				"0001",
				"9000",
				"0900",
				"0090",
				"0009",
			},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			out := neighbours(tc.in, tc.level)

			for _, outCombination := range tc.out {
				found := false
				for _, neighbour := range out {
					if outCombination == neighbour.node {
						found = true
					}
				}
				if !found {
					t.Errorf("neighbour %v not found in %v", outCombination, out)
				}
			}
		})
	}
}

func TestOpenLock(t *testing.T) {
	tests := []struct {
		deadends []string
		target   string
		level    int
	}{
		{
			deadends: []string{"0201", "0101", "0102", "1212", "2002"},
			target:   "0202",
			level:    6,
		},
		{
			deadends: []string{"8888"},
			target:   "0009",
			level:    1,
		},
		{
			deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
			target:   "8888",
			level:    -1,
		},
		{
			deadends: []string{"5557", "5553", "5575", "5535", "5755", "5355", "7555", "3555", "6655", "6455", "4655", "4455", "5665", "5445", "5645", "5465", "5566", "5544", "5564", "5546", "6565", "4545", "6545", "4565", "5656", "5454", "5654", "5456", "6556", "4554", "4556", "6554"},
			target:   "5555",
			level:    -1,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			level := openLock(tc.deadends, tc.target)
			if level != tc.level {
				t.Errorf("expected level %d got %d", tc.level, level)
			}
		})
	}
}
