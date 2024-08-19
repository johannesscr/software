package challenges

/*
# 650. 2 Keys Keyboard
> Medium

_Hint_
There is only one character 'A' on the screen of a notepad. You can perform one
of two operations on this notepad for each step:

Copy All: You can copy all the characters present on the screen (a partial copy
is not allowed).

Paste: You can paste the characters which are copied last time.

Given an integer n, return the minimum number of operations to get the character
'A' exactly n times on the screen.

---

Example 1:

Input: n = 3
Output: 3
Explanation: Initially, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.

---

Example 2:

Input: n = 1
Output: 0

---

Constraints:

1 <= n <= 1000
*/

/*
APPROACH I
To get use copy all and paste to get to the nearest square root of a number.
Then use the square root the same number of times to get to the number.

Example n = 100
Step 1 copy all 1A or "A".
Step 2-10 paste 1A to get "AAAAAAAAAA" or 10A.
Then copy 10A.
Paste 10 times to get 100A.
*/

func minSteps(n int) int {
	return 0
}
