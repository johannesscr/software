package challenges

import (
	"fmt"
	"strconv"
	"testing"
)

/*
You are given an array of strings tokens that represents an arithmetic
expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the
expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit
integer.


Example 1:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Example 2:

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Example 3:

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

Constraints:

1 <= tokens.length <= 104
tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
*/

/*
We push all the operands onto the stack, when we encounter a value, we check if
the top of the stack is also a value if it is, we pop the top two off the stack
as it will be a value followed by an operand. We perform the calculation and
push it back onto the stack and continue.

since we are working with an array, it is inherently a stack, we just step
through the array/stack in reverse.
*/

func evalRPN2(tokens []string) int {
	if len(tokens) < 3 {
		val, _ := strconv.ParseInt(tokens[0], 10, 0)
		return int(val)
	}
	for len(tokens) >= 3 {
		// pop from the stack
		//left := tokens[0]
		var calc int64 = 0
		left, _ := strconv.ParseInt(tokens[0], 10, 0)
		right, _ := strconv.ParseInt(tokens[1], 10, 0)
		operand := tokens[2]
		switch operand {
		case "+":
			calc = left + right
		case "-":
			calc = left - right
		case "*":
			calc = left * right
		case "/":
			calc = left / right
		}
		tokens = append([]string{fmt.Sprintf("%d", calc)}, tokens[3:]...)
	}
	total, _ := strconv.ParseInt(tokens[0], 10, 0)
	return int(total)
}

/*
we step through the array of tokens and add each token to the stack of values
when we encounter an operand, the top two values in the stack are popped off
the value is calculated and the value is push onto the stack. This process
is continued until we have stepped through the entire array.
*/

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			stackLen := len(stack)
			l := stack[stackLen-2]
			r := stack[stackLen-1]
			stack = append(stack[:stackLen-2], l+r)
		case "-":
			stackLen := len(stack)
			l := stack[stackLen-2]
			r := stack[stackLen-1]
			stack = append(stack[:stackLen-2], l-r)
		case "*":
			stackLen := len(stack)
			l := stack[stackLen-2]
			r := stack[stackLen-1]
			stack = append(stack[:stackLen-2], l*r)
		case "/":
			stackLen := len(stack)
			l := stack[stackLen-2]
			r := stack[stackLen-1]
			stack = append(stack[:stackLen-2], l/r)
		default:
			v, _ := strconv.ParseInt(tokens[i], 10, 0)
			stack = append(stack, int(v))
		}
	}
	return stack[0]
}

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{
			in:  []string{"2", "1", "+", "3", "*"},
			out: 9,
		},
		{
			in:  []string{"4", "13", "5", "/", "+"},
			out: 6,
		},
		{
			in:  []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			out: 22,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			out := evalRPN(tc.in)
			if out != tc.out {
				t.Errorf("expected value %d got %d", tc.out, out)
			}
		})
	}
}
