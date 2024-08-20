package array

import (
	"testing"
)

/*
# 1089. Duplicate Zeros
> Easy
_Hint_
Given a fixed-length integer array arr, duplicate each occurrence of zero,
shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. Do
the above modifications to the input array in place and do not return anything.

---

Example 1:

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to:
[1,0,0,2,3,0,0,4]

---

Example 2:

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to:
[1,2,3]

---

Constraints:

1 <= arr.length <= 104
0 <= arr[i] <= 9
*/

func TestDuplicateZeros(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		out  []int
	}{
		{
			name: "case one",
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			out:  []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "case two",
			arr:  []int{1, 2, 3},
			out:  []int{1, 2, 3},
		},
		{
			name: "case two",
			arr:  []int{8, 4, 5, 0, 0, 0, 0, 7},
			out:  []int{8, 4, 5, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			duplicateZeros(test.arr)
			// fmt.Println(test.arr)
		})
	}
}

/*
APPROACH I

step through the array
when a zero is encoutered start tracking the number of subsequent zeros
when a non-zero is encountered the sequence ends and to the inplace update.
when the update is complete reset the counters until the next zero is encoutered
or the end of the array has been reached.
*/
func duplicateZeros(arr []int) {
	numZeros := 0
	for i := 0; i < len(arr); i++ {
		// fmt.Println(i, arr, arr[i], numZeros)
		if arr[i] == 0 {
			// we have encountered a zero
			numZeros++
		} else {
			// we are on the first non-zero instance after encoutering a zero
			// we have a non-zero
			if numZeros > 0 {
				// we need to insert the new zeros
				// do a single step offset to shift all values
				// start from the back with a num zero offset and shift all
				// values to the left by the num zero offset until i + offset is
				// reached
				for j := len(arr) - 1; j >= i+numZeros; j-- {
					// fmt.Println("j -", j)
					arr[j] = arr[j-numZeros]
				}
				// we are on the first non-zero
				// replace all the values with zeros
				for j := 0; j < numZeros; j++ {
					// check that j and the offset do not exceed the length
					// of the array
					// fmt.Println("j+", j, i+j, arr)
					if j >= len(arr) || i+j >= len(arr) {
						break
					}
					arr[i+j] = 0
				}
				// skip forward by the number of zeros that we have inserted
				i = i + numZeros
				// reset the zero counter
				numZeros = 0
			}
		}
	}
}

/*
APPROACH II
Fastest
*/

func duplicateZerosII(arr []int) {
	ma := make([]int, len(arr))
	i := 0
	for _, n := range arr {
		if i >= len(arr) {
			break
		}
		if n != 0 {
			ma[i] = n
		} else {
			i++
		}
		i++
	}
	for i, n := range ma {
		arr[i] = n
	}
}
