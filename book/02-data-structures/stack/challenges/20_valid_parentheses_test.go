package challenges

import (
	"strings"
	"testing"
)

/*
**Valid Parentheses**

Given a string `s` containing just the characters `'(', ')', '{', '}', '['` and
`']'`, determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

/*
**Notes**
1. If there is an odd number of characters, then the string will automatically
   be invalid.
2. Add all open brackets to a stack.
3. When you encounter a closing bracket, remove the top element from the stack.
4. If the "popped" element matches continue, else the string is invalid.
*/

func isValid(s string) bool {
	// we have an odd number of characters so it will default to false
	if len(s)%2 != 0 {
		return false
	}

	// create out stack
	stack := make([]string, 0)
	brackets := strings.Split(s, "")

	for _, bracket := range brackets {
		switch bracket {
		case "(", "{", "[":
			// open bracket push onto the stack
			stack = append(stack, bracket)
		case ")":
			// peek if the top of the stack contains the other
			if len(stack) == 0 {
				return false
			}
			openBracket := stack[len(stack)-1]
			if openBracket == "(" {
				// pop off from the stack
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "}":
			// peek if the top of the stack contains the other
			if len(stack) == 0 {
				return false
			}
			openBracket := stack[len(stack)-1]
			if openBracket == "{" {
				// pop off from the stack
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "]":
			// peek if the top of the stack contains the other
			if len(stack) == 0 {
				return false
			}
			openBracket := stack[len(stack)-1]
			if openBracket == "[" {
				// pop off from the stack
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		default:
			return false
		}
	}

	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "()",
			out: true,
		},
		{
			in:  "()[]{}",
			out: true,
		},
		{
			in:  "(]",
			out: false,
		},
		{
			in:  "(()",
			out: false,
		},
		{
			in:  "((",
			out: false,
		},
		{
			in:  "){",
			out: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			out := isValid(tc.in)
			if out != tc.out {
				t.Errorf("expected %t got %t", tc.out, out)
			}
		})
	}
}

/* V2 */

type Stack struct {
	buffer []byte
}

func (s *Stack) push(b byte) {
	s.buffer = append(s.buffer, b)
}

func (s *Stack) pop() byte {
	if len(s.buffer) == 0 {
		panic("Trying to pop from empty stack")
	}

	result := s.buffer[len(s.buffer)-1]
	s.buffer = s.buffer[:len(s.buffer)-1]

	return result
}

func (s *Stack) peek() byte {
	if len(s.buffer) == 0 {
		panic("Trying to pop in empty stack")
	}
	return s.buffer[len(s.buffer)-1]
}

func (s *Stack) empty() bool {
	return len(s.buffer) == 0
}

func symmetrical(left byte, right byte) bool {
	if left == '{' && right == '}' {
		return true
	}
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}

	return false
}

func isValid2(s string) bool {
	stack := Stack{buffer: make([]byte, 0)}

	length := len(s)

	if length%2 == 1 {
		return false
	}

	for i := 0; i < length; i += 1 {
		if !stack.empty() {
			last := stack.peek()

			if symmetrical(last, s[i]) {
				stack.pop()
			} else {
				stack.push(s[i])
			}
		} else {
			stack.push(s[i])
		}

	}

	return stack.empty()
}
