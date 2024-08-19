package challenges

import (
	"testing"
)

/*
# 2053. Kth Distinct String in an Array
> Easy
_Hint_
A distinct string is a string that is present only once in an array.

Given an array of strings arr, and an integer k, return the kth distinct string
present in arr. If there are fewer than k distinct strings, return an empty
string "".

Note that the strings are considered in the order in which they appear in the
array.

---

Example 1:

Input: arr = ["d","b","c","b","c","a"], k = 2
Output: "a"
Explanation:
The only distinct strings in arr are "d" and "a".
"d" appears 1st, so it is the 1st distinct string.
"a" appears 2nd, so it is the 2nd distinct string.
Since k == 2, "a" is returned.

---

Example 2:

Input: arr = ["aaa","aa","a"], k = 1
Output: "aaa"
Explanation:
All strings in arr are distinct, so the 1st string "aaa" is returned.

---

Example 3:

Input: arr = ["a","b","a"], k = 3
Output: ""
Explanation:
The only distinct string is "b". Since there are fewer than 3 distinct strings, we return an empty string "".

---

Constraints:

1 <= k <= arr.length <= 1000
1 <= arr[i].length <= 5
arr[i] consists of lowercase English letters.
*/

func TestKthDistinct(t *testing.T) {
	tests := []struct {
		name   string
		arr    []string
		k      int
		output string
	}{
		{
			name:   "case one",
			arr:    []string{"d", "b", "c", "b", "c", "a"},
			k:      2,
			output: "a",
		},
		{
			name:   "case two",
			arr:    []string{"aaa", "aa", "a"},
			k:      1,
			output: "aaa",
		},
		{
			name:   "case three",
			arr:    []string{"a", "b", "a"},
			k:      3,
			output: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			o := kthDistinctI(test.arr, test.k)
			if test.output != o {
				t.Errorf("expect %s got %s", test.output, o)
			}
		})
	}
}

/*
APPROACH I:
A distinct string only appears once, that is ["b", "b"] would contain no
distinct string, whereas, ["b", "b", "a"] would contain "a" as the only distinct
string and ["a", "b"] would contain two "a" and "b" distinct strings.

Simple solution:
Keep a slice of distinct strings in the order in which they are observed.
Step through each string in the array.
- Check if the string exists in the slice of distinct string.
- If not, add it to the list of distinct string.
- If it does, then remove the current string from the list of distinct strings.
Finally, do the k-th lookup of a distinct string.
*/

func kthDistinctI(arr []string, k int) string {
	xd := make([]string, 0)
	seen := make(map[string]bool)
	for i := 0; i < len(arr); i++ {
		s := arr[i]
		for j, ds := range xd {
			if s == ds {
				xd = append(xd[:j], xd[j+1:]...)
				seen[s] = true
				break
			}
		}
		if seen[s] == false {
			xd = append(xd, s)
		}
	}
	if k <= len(xd) {
		return xd[k-1]
	}
	return ""
}

/*
APPROACH II
FASTEST
*/
func kthDistinctII(arr []string, k int) string {
	mp := make(map[string]int)
	for _, val := range arr {
		mp[val] += 1
	}
	ku := 1

	for _, val := range arr {
		if mp[val] == 1 && k == ku {
			return val
		} else if mp[val] == 1 {
			ku++
		}
	}
	return ""
}
