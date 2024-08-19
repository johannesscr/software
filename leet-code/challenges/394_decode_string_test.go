package challenges

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
394. Decode String

Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the
square brackets is being repeated exactly k times. Note that k is guaranteed to
be a positive integer.

You may assume that the input string is always valid; there are no extra white
spaces, square brackets are well-formed, etc. Furthermore, you may assume that
the original data does not contain any digits and that digits are only for
those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed
105.

Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"
Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"
Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"

Constraints:

1 <= s.length <= 30
s consists of lowercase English letters, digits, and square brackets '[]'.
s is guaranteed to be a valid input.
All the integers in s are in the range [1, 300].
*/

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "case 1",
			input:  "3[a]2[bc]",
			output: "aaabcbc",
		},
		{
			name:   "case 2",
			input:  "3[a2[c]]",
			output: "accaccacc",
		},
		{
			name:   "case 3",
			input:  "2[abc]3[cd]ef",
			output: "abcabccdcdcdef",
		},
		{
			name:   "case 4",
			input:  "10[leetcode]",
			output: "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			o := decodeString(tc.input)
			if o != tc.output {
				t.Errorf("expected %s => %s got %s", tc.input, tc.output, o)
			}
		})
	}
}

/*
Solution 1:
create a stack, when you reach "]" pop off the stack until you get "[", pop
another off to get the repetition number, repeat the string and put it back
onto the stack.
*/

func decodeString(s string) string {
	stack := make([]rune, 0)

	//loop := true
	stringIndex := 0
	parse := false

	for {
		if !parse {
			r := rune(s[stringIndex])
			if r == ']' {
				parse = true
			} else {
				stack = append(stack, r)
			}
			stringIndex++
		}
		//fmt.Println("\nstringIndex", stringIndex, len(s))
		//fmt.Println("parse", parse)
		//fmt.Println("stack", stack, string(stack))
		if parse {
			parseStack := make([]rune, 0)
			for i := len(stack) - 1; i > 0; i-- {
				// get the character or rune
				r := stack[i]
				// pop off the stack
				stack = stack[:len(stack)-1]
				//fmt.Println("\tr", r, string(r))
				//fmt.Println("\tstack", stack, string(stack))
				//fmt.Println("\tparsestack", parseStack, string(parseStack))
				if r == '[' {
					repeatNum := 0
					numStackReverse := make([]rune, 0)
					for j := i - 1; j >= 0; j-- {
						numStack := []rune{stack[j]}
						numStack = append(numStack, numStackReverse...)
						num, err := strconv.Atoi(string(numStack))
						if err != nil {
							// we went too far
							break
						} else {
							repeatNum = num
							numStackReverse = numStack
							// shorten the stack
							stack = stack[:j]
						}
						fmt.Println("numStack", string(numStack))
					}
					// pop the number off the stack
					//stack = stack[:len(stack)-1]

					repeated := strings.Repeat(string(parseStack), repeatNum)
					for _, re := range repeated {
						stack = append(stack, re)
					}
					break
				} else {

					newParseStack := []rune{r}
					parseStack = append(newParseStack, parseStack...)
				}
			}
			parse = false
		}
		if stringIndex == len(s) {
			break
		}
	}

	b := string(stack)
	//fmt.Println(b)
	return b
}
