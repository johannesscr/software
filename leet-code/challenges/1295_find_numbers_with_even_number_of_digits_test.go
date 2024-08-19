package challenges

import (
	"strconv"
	"testing"
)

/*
# 1295. Find Numbers with Even Number of Digits
> Easy
_Hint_
Given an array nums of integers, return how many of them contain an even number of digits.

---

Example 1:

Input: nums = [12,345,2,6,7896]
Output: 2
Explanation:
12 contains 2 digits (even number of digits).
345 contains 3 digits (odd number of digits).
2 contains 1 digit (odd number of digits).
6 contains 1 digit (odd number of digits).
7896 contains 4 digits (even number of digits).
Therefore only 12 and 7896 contain an even number of digits.

---

Example 2:

Input: nums = [555,901,482,1771]
Output: 1
Explanation:
Only 1771 contains an even number of digits.

---

Constraints:

1 <= nums.length <= 500
1 <= nums[i] <= 105
*/

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "case one",
			input:  []int{12, 345, 2, 6, 7896},
			output: 2,
		},
		{
			name:   "case two",
			input:  []int{555, 901, 482, 1771},
			output: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := findNumbers(test.input)
			if test.output != o {
				t.Errorf("expected %d got %d", test.output, o)
			}
		})
	}
}

func findNumbers(nums []int) int {
	x := 0
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			x++
		}
	}
	return x
}
