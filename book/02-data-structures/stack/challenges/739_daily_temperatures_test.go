package challenges

import (
	"fmt"
	"math"
	"testing"
)

/*
**739 Daily Temperatures**

Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to
wait after the ith day to get a warmer temperature. If there is no future day
for which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]

Constraints:

`1 <= temperatures.length <= 105`
`30 <= temperatures[i] <= 100`
*/

/*
Initialise the answers array to be the same length as the number of temperatures
to be compared.

Iterate through the temperatures at each index.
- check the top temperature of the stack.
	- **if** the top temperature is greater than the current temperature, then
	  the temperature is decreasing and add the current index and temperature
	  to the stack.
	- **else** the current temperature is greater than the temperature on top
	  of the stack. Therefore, pop the top temperature off the stack and update
	  its answer to be the difference in indices between the current
	  temperature and the one just popped off. Check this for all temperature
	  on the stack **iff** the current temperature is greater that the
	  temperature on the stack.
*/

// stack is lower case as we do not want to export it
type item struct {
	index int
	value int
}

type stack []item

func (s stack) push(i item) stack {
	return append(s, i)
}

func (s stack) pop() (*item, stack) {
	if s.empty() {
		return nil, s
	}
	lastIndex := len(s) - 1
	return &s[lastIndex], s[:lastIndex]
}

func (s stack) top() int {
	if s.empty() {
		return math.MaxInt
	}
	return s[len(s)-1].value
}

func (s stack) empty() bool {
	return len(s) == 0
}

// dailyTemperatures uses a monotonic stack.
func dailyTemperatures(temperatures []int) []int {
	answers := make([]int, len(temperatures))
	st := stack{}

	for i, temp := range temperatures {
		// only if the top temperature is greater than the current
		for temp > st.top() {
			prev, newStack := st.pop()
			answers[prev.index] = i - prev.index
			st = newStack
		}
		if st.empty() || st.top() >= temp {
			st = st.push(item{i, temp})
		}
	}

	return answers
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{73, 74, 75, 71, 69, 72, 76, 73},
			output: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			input:  []int{30, 40, 50, 60},
			output: []int{1, 1, 1, 0},
		},
		{
			input:  []int{30, 60, 90},
			output: []int{1, 1, 0},
		},
		{
			input:  []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70},
			output: []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			out := dailyTemperatures(tc.input)
			s1 := fmt.Sprintf("%v", out)
			s2 := fmt.Sprintf("%v", tc.output)
			if s1 != s2 {
				t.Errorf("expected answers %s got %s", s2, s1)
			}
		})
	}
}
