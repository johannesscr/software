package challenges

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{
			in:  121,
			out: true,
		},
		{
			in:  -121,
			out: false,
		},
		{
			in:  12321,
			out: true,
		},
		{
			in:  1221,
			out: true,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			o := IsPalindrome(tc.in)
			if o != tc.out {
				t.Errorf("expected %t for %d got %t", tc.out, tc.in, o)
			}
		})
	}
}

/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Definition:
*Palindrome*: An integer is a palindrome when it reads the same forward and
backward. For example, 121 is a palindrome while 123 is not.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it
becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:
- `-2^31 <= x <= 2^31 - 1`
*/

func IsPalindrome(x int) bool {
	// convert the integer to a string
	left := fmt.Sprintf("%d", x)
	right := ""
	for i := len(left) - 1; i >= 0; i-- {
		right = fmt.Sprintf("%s%c", right, left[i])
	}
	return left == right
}
