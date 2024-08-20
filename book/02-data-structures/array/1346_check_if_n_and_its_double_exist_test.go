package array

import (
	"testing"
)

/*
# 1346. Check If N and Its Double Exist
> Easy

_Hint_
Given an array arr of integers, check if there exist two indices i and j such
that:

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

---

Example 1:

Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

---

Example 2:

Input: arr = [3,1,7,11]
Output: false
Explanation: There is no i and j that satisfy the conditions.

---

Constraints:

2 <= arr.length <= 500
-103 <= arr[i] <= 103
*/

func TestCheckIfExist(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		out  bool
	}{
		{
			name: "case one",
			arr:  []int{10, 2, 5, 3},
			out:  true,
		},
		{
			name: "case two",
			arr:  []int{3, 1, 7, 11},
			out:  false,
		},
		{
			name: "case three",
			arr:  []int{7, 1, 14, 11},
			out:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := checkIfExist(test.arr)
			if o != test.out {
				t.Errorf("expected %t got %t", test.out, o)
			}
		})
	}
}

/*
APPROACH I
*/

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == 2*arr[j] || arr[j] == 2*arr[i] {
				return true
			}
		}
	}
	return false
}

/*
APPROACH II
Fastest
*/

func checkIfExistII(arr []int) bool {
	set := make(map[int]int, len(arr))

	// Swap keys and values
	for i := range arr {
		set[arr[i]] = i
	}

	for i := range arr {
		// check each i if there is a double in the map
		// and that i is not the same index as j.
		if j, ok := set[arr[i]*2]; ok && i != j {
			return true
		}
	}

	return false
}
