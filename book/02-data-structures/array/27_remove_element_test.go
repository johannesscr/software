package array

import (
	"fmt"
	"testing"
)

/*
# 27. Remove Element
> Easy

_Hint_
Given an integer array nums and an integer val, remove all occurrences of val in
nums in-place. The order of the elements may be changed. Then return the number
of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get
accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contains the
elements which are not equal to val. The remaining elements of nums are not
important as well as the size of nums.

Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                            // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

---

Example 1:

Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

---

Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).

---

Constraints:

0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100
*/

func TestRemoveElemet(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		k    int
		out  []int
	}{
		{
			name: "case one",
			nums: []int{3, 2, 2, 3},
			val:  3,
			k:    2,
			out:  []int{2, 2},
		},
		{
			name: "case two",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			k:    5,
			out:  []int{0, 1, 3, 0, 4},
		},
		{
			name: "case three",
			nums: []int{},
			val:  3,
			k:    0,
			out:  []int{},
		},
		{
			name: "case four",
			nums: []int{2},
			val:  3,
			k:    1,
			out:  []int{2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			k := removeElement(test.nums, test.val)
			if k != test.k {
				t.Errorf("expected k %d got %d", test.k, k)
			}

			s1 := fmt.Sprintf("%v", test.out)
			s2 := fmt.Sprintf("%v", test.nums[:k])
			if s1 != s2 {
				t.Errorf("expected nums %v got %v", test.out, test.nums[:k])
			}
		})
	}
}

/*
APPROACH I

Loop through the array.
If the current index value equals "val" do a shift operation for that index.
- remember to decrement the counter, to re-evaluate after the shift.

This will not work, as you will reach the end and continuously do a
re-evaluation.
*/

/*
APPROACH II

Work backwards, identify the number of "val" matches.
Loop and do a left shift for the total count of matches.
*/

/*
APPROACH III

Create a duplicate array. Initialise with -1 based on the constraint of [0,50].
- Keep track of length as this is k.
Loop through nums:
- If nums[i] is val, then skip.
- Else add nums[i] at duplicate length to duplicate.
Map the duplicate back to nums.
*/

func removeElement(nums []int, val int) int {
	d := make([]int, len(nums))
	k := 0

	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			// skipp
		} else {
			d[k] = nums[i]
			k++
		}
	}

	for j := 0; j < k; j++ {
		nums[j] = d[j]
	}

	return k
}

/*
APPROACH IV
> Fastest

We replace by offset, as i steps through the entire array, j is just the offset
skipping all the values equal to val.
*/

func removeElementIV(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
