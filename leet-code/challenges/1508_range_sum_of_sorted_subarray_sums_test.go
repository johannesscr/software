package challenges

import (
	"slices"
	"sort"
	"testing"
)

/*
# 1508. Range Sum of Sorted Subarray Sums
> Medium

_Hint_
You are given the array nums consisting of n positive integers. You computed the
sum of all non-empty continuous subarrays from the array and then sorted them in
non-decreasing order, creating a new array of n * (n + 1) / 2 numbers.

Return the sum of the numbers from index left to index right (indexed from 1),
inclusive, in the new array. Since the answer can be a huge number return it
modulo 10^9 + 7.

---

Example 1:

Input: nums = [1,2,3,4], n = 4, left = 1, right = 5
Output: 13
Explanation:
All subarray sums are 1, 3, 6, 10, 2, 5, 9, 3, 7, 4. After sorting them in
non-decreasing order we have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. The
sum of the numbers from index le = 1 to ri = 5 is 1 + 2 + 3 + 3 + 4 = 13.

---

Example 2:

Input: nums = [1,2,3,4], n = 4, left = 3, right = 4
Output: 6
Explanation: The given array is the same as example 1. We have the new array
[1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. The sum of the numbers from index le = 3 to
ri = 4 is 3 + 3 = 6.

---

Example 3:

Input: nums = [1,2,3,4], n = 4, left = 1, right = 10
Output: 50

---

Constraints:

n == nums.length
1 <= nums.length <= 1000
1 <= nums[i] <= 100
1 <= left <= right <= n * (n + 1) / 2
*/

func aOne() []int {
	a := make([]int, 1000)
	for i := 0; i < len(a); i++ {
		a[i] = 100
	}
	return a
}

const MODULUS = 1000000000 + 7
const MODULUS1 = 10e8 + 7

func TestRangeSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		n      int
		left   int
		right  int
		output int
	}{
		{
			name:   "case one",
			nums:   []int{1, 2, 3, 4},
			n:      4,
			left:   1,
			right:  5,
			output: 13,
		},
		{
			name:   "case two",
			nums:   []int{1, 2, 3, 4},
			n:      4,
			left:   3,
			right:  4,
			output: 6,
		},
		{
			name:   "case three",
			nums:   []int{1, 2, 3, 4},
			n:      4,
			left:   1,
			right:  10,
			output: 50,
		},
		{
			name:   "case four",
			nums:   aOne(),
			n:      1000,
			left:   1,
			right:  500500,
			output: 716699888,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := rangeSumII(test.nums, test.n, test.left, test.right)
			if o != test.output {
				t.Errorf("expected %d got %d", test.output, o)
			}
		})
	}
}

/*
APPROACH I
Brute force: Break down the steps.
1. Loop through from [i, end) of the array.
2. For each range, create the new subarray.
	- can do it with a loop and append to an array.
	- can do it with recursion and return the result. Can get quite large with
      the constraint of length 1000.
3. Sort the full array.
4. Sum the range between left and right.
*/

func rangeSumI(nums []int, n int, left int, right int) int {
	expanded := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		subarray := []int{nums[i]}
		offset := i
		for j := 1; j < n-offset; j++ {
			subarray = append(subarray, subarray[j-1]+nums[offset+j])
		}
		expanded = append(expanded, subarray...)
	}
	// sort in place, slices.sort is slightly faster than sort.sort.
	slices.Sort(expanded)
	sum := 0
	for _, val := range expanded[left-1 : right] {
		sum += val
	}
	return sum % (10e8 + 7)
}

func rangeSumII(nums []int, n int, left int, right int) int {
	expanded := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		subarray := []int{sum}
		for j := i + 1; j < n; j++ {
			sum = sum + nums[j]
			subarray = append(subarray, sum)
		}
		expanded = append(expanded, subarray...)
	}
	// sort in place, slices.sort is slightly faster than sort.sort.
	slices.Sort(expanded)
	sum := 0
	for _, val := range expanded[left-1 : right] {
		sum += val
	}
	return sum % (10e8 + 7)
}

func rangeSumIII(nums []int, n int, left int, right int) int {
	expanded := make([]int, 0)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum = sum + nums[j]
			expanded = append(expanded, sum)
		}
	}
	// sort in place, slices.sort is slightly faster than sort.sort.
	slices.Sort(expanded)
	sum := 0
	for _, val := range expanded[left-1 : right] {
		sum += val
	}
	return sum % (10e8 + 7)
}

func rangeSum(nums []int, n int, left int, right int) int {
	const MOD = 1000000007
	subarraySums := []int{}

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			subarraySums = append(subarraySums, sum)
		}
	}

	sort.Ints(subarraySums)

	sum := 0
	for i := left - 1; i < right; i++ {
		sum = (sum + subarraySums[i]) % MOD
	}

	return sum
}

/*
APPROACH IIII: Sliding Window
*/

func rangeSumIIII(nums []int, n int, left int, right int) int {
	return 0
}
