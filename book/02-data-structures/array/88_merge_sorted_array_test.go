package array

import (
	"fmt"
	"testing"
)

/*
# 88. Merge Sorted Array
> Easy

_Hint_
You are given two integer arrays nums1 and nums2, sorted in non-decreasing
order, and two integers m and n, representing the number of elements in nums1
and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be
stored inside the array nums1. To accommodate this, nums1 has a length of m + n,
where the first m elements denote the elements that should be merged, and the
last n elements are set to 0 and should be ignored. nums2 has a length of n.

---

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

---

Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].

---

Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

---

Constraints:

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109

---

Follow up: Can you come up with an algorithm that runs in O(m + n) time?
*/

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		out   []int
	}{
		{
			name:  "case one",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			out:   []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "case two",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			out:   []int{1},
		},
		{
			name:  "case four",
			nums1: []int{2, 0},
			m:     1,
			nums2: []int{1},
			n:     1,
			out:   []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mergeII(test.nums1, test.m, test.nums2, test.n)
			fmt.Println(test.nums1)
		})
	}
}

/*
APPROACH I
create a new array of with capacity n + m
step through both arrays a total of n + m times
keep counters i and j for each array

for each step
check which array index has the smallest number add that number to the new array
and increment the index.
*/

func mergeI(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, n+m)
	i, j := 0, 0
	if len(nums2) == 0 {
		return
	}
	// accommodate for the two zero indices
	for k, _ := range nums {
		if i < m && j < m && nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			if i < m {
				i++
			}
		} else {
			nums[k] = nums2[j]
			if j < n {
				j++
			}
		}
	}
	// nums1 = nums
	for k, _ := range nums1 {
		nums1[k] = nums[k]
	}
}

func mergeII(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, n+m)
	i, j := 0, 0
	if len(nums2) == 0 {
		return
	}
	// accommodate for the two zero indices
	for k, _ := range nums {
		n1, ok1 := get(nums1, i, m)
		n2, ok2 := get(nums2, j, n)

		if ok1 && ok2 {
			if n1 < n2 {
				nums[k] = n1
				i++
			} else {
				nums[k] = n2
				j++
			}
		} else if ok1 {
			nums[k] = n1
			i++
		} else {
			nums[k] = n2
			j++
		}
	}
	// nums1 = nums
	for k, _ := range nums1 {
		nums1[k] = nums[k]
	}
}

func get(nums []int, index int, len int) (int, bool) {
	if index > len-1 {
		return 0, false
	}
	return nums[index], true
}

/*
APPROACH III
Work Backwards
*/

func mergeIII(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, n+m)
	i, j := m-1, n-1

	for k := n + m - 1; k >= 0; k-- {
		if i > -1 && j > -1 {
			if nums1[i] > nums2[j] {
				nums[k] = nums1[i]
				i--
			} else {
				nums[k] = nums2[j]
				j--
			}
		} else if i > -1 {
			nums[k] = nums1[i]
			i--
		} else if j > -1 {
			nums[k] = nums2[j]
			j--
		}
	}
	for k, _ := range nums1 {
		nums1[k] = nums[k]
	}
}

/*
APPROACH IV
Fastest
Backwards
Decrement the length param of each input array
*/
func mergeIIII(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	m -= 1
	n -= 1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m -= 1
		} else {
			nums1[i] = nums2[n]
			n -= 1
		}
		i -= 1
	}
	for m >= 0 {
		nums1[i] = nums1[m]
		i -= 1
		m -= 1
	}
	for n >= 0 {
		nums1[i] = nums2[n]
		i -= 1
		n -= 1
	}
}
