package array

import (
	"fmt"
	"testing"
)

/*
# 26. Remove Duplicates from Sorted Array
> Easy

_Hint_
Given an integer array nums sorted in non-decreasing order, remove the
duplicates in-place such that each unique element appears only once. The
relative order of the elements should be kept the same. Then return the number
of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you
need to do the following things:

Change the array nums such that the first k elements of nums contain the unique
elements in the order they were present in nums initially. The remaining
elements of nums are not important as well as the size of nums.

Return k.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.

---

Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

---

Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

---

Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.
*/

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		out  []int
	}{
		{
			name: "case one",
			nums: []int{1, 1, 2},
			k:    2,
			out:  []int{1, 2},
		},
		{
			name: "case two",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			k:    5,
			out:  []int{0, 1, 2, 3, 4},
		},
		{
			name: "case three",
			nums: []int{},
			k:    0,
			out:  []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			k := removeDuplicates(test.nums)
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
keep a map of all seen integers, keep a j offset.
if number has already been seen set j offset to num i and update the map.
*/

func removeDuplicates(nums []int) int {
	unique := make(map[int]bool, 0)
	j := 0

	for i := 0; i < len(nums); i++ {
		if unique[nums[i]] {
			// we have seen this value
		} else {
			// new unique value
			unique[nums[i]] = true
			nums[j] = nums[i]
			j++
		}
	}
	return len(unique)
}

/*
APPROACH II
Check the previous as the array is sorted
*/

func removeDuplicatesII(nums []int) int {
	dest := 1
	k := 1
	previous := nums[0]
	for i := 1; i < len(nums); i++ {
		if previous != nums[i] {
			nums[dest] = nums[i]
			dest++
			k++
			previous = nums[i]
		}
	}
	return k
}
