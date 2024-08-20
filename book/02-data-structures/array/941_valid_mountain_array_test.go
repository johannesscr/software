package array

import (
	"testing"
)

/*
# 941. Valid Mountain Array
> Easy

_Hint_
Given an array of integers arr, return true if and only if it is a valid
mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

---

Example 1:

Input: arr = [2,1]
Output: false

---

Example 2:

Input: arr = [3,5,5]
Output: false

---

Example 3:

Input: arr = [0,3,2,1]
Output: true

---

Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 104
*/

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		out  bool
	}{
		{
			name: "case one",
			arr:  []int{2, 1},
			out:  false,
		},
		{
			name: "case two",
			arr:  []int{3, 5, 5},
			out:  false,
		},
		{
			name: "case three",
			arr:  []int{0, 3, 2, 1},
			out:  true,
		},
		{
			name: "case four",
			arr:  []int{0, 1, 5, 2, 4},
			out:  false,
		},
		{
			name: "case five",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:  false,
		},
		{
			name: "case six",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			out:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := validMountainArray(test.arr)
			if o != test.out {
				t.Errorf("expected %t got %t", test.out, o)
			}
		})
	}
}

/*
APPROACH I

We will use a sliding window with two states UP and DOWN.
In either state if arr[i] == arr[j] it is not a strict mountain => false
In UP arr[i] < arr[j] where i < j
- If we encounter arr[i] > arr[j] switch state to DOWN
In DOWN arr[i] > arr[j] where i < j
*/

const (
	UP = iota
	DOWN
)

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	if arr[1] < arr[0] {
		// the mountain must start by going up
		return false
	}

	going := UP
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			return false
		}

		switch going {
		case UP:
			if arr[i-1] > arr[i] {
				going = DOWN
			}
			break
		case DOWN:
			if arr[i-1] < arr[i] {
				return false
			}
			break
		}
	}

	// check that we indeed did go down
	return going == DOWN
}

/*
APPROACH II

We take a step as a sliding window, we assume we are going up all the way
until the top. At the first point we start going down, we should only be
going down from there onwards. If we start going up again it is not a valid
mountain.
*/

func validMountainArrayII(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	if arr[1] < arr[0] {
		// the mountain must start by going up
		return false
	}

	increase := true

	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			return false
		}

		if increase {
			// we have reached the top and are now going down
			if arr[i-1] > arr[i] {
				increase = false
			}
		} else {
			if arr[i-1] < arr[i] {
				return false
			}
		}
	}

	// check that we indeed did go down
	return increase == false
}
