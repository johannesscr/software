package challenges

import (
	"math"
	"testing"
)

/*
Given an integer n, return the least number of perfect square numbers that sum
to n.

A perfect square is an integer that is the square of an integer; in other words,
it is the product of some integer with itself. For example, 1, 4, 9, and 16 are
perfect squares while 3 and 11 are not.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.

Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.

Constraints:

1 <= n <= 104
*/

/* ATTEMPT ONE */

func nextLargestSquare(n int) int {
	x := math.Sqrt(float64(n))
	return int(math.Floor(x))
}

// NumSquares returns the least number of perfect square numbers that sum to n.
// By using the largest square number, we first find all the multiples for the
// largest square number in reducing order.
//
// Discovered that this is not always optimal, as using a smaller square number
// may result in a smaller number of squares.
func NumSquaresV1(n int) int {
	numSq := 0
	remainder := n
	for remainder > 0 {
		sq := nextLargestSquare(remainder)
		num := sq * sq
		multiples := remainder / num
		remainder %= multiples * num
		numSq += multiples
	}
	return numSq
}

func TestNumSquares(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		// failed here
		//{
		//	n:        12,
		//	expected: 3,
		//},
		{
			n:        13,
			expected: 2,
		},
	}
	for _, test := range tests {
		actual := NumSquaresV1(test.n)
		if actual != test.expected {
			t.Errorf("NumSquares(%d) = %d; expected %d", test.n, actual, test.expected)
		}
	}
}

/* RECURSION */

// NumSquares returns the least number of perfect square numbers that sum to n.
//
// To do this it uses a recursive method of taking the current number, and for
// each square number less than the current number, it subtracts the square
// number from the current number. It then checks the number of squares the
// previous number has and adds one to it. It then takes the minimum of the
// current number of squares and the previous number of squares.
func NumSquaresV2(n int) int {
	// find all the square numbers less than n
	var squares []int
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	numbers := make([]int, n+1)
	numbers[0] = 0

	// incrementally find the number of squares for each number up to n
	for i := 1; i <= n; i++ {
		// set the current number of squares to the max int
		numbers[i] = math.MaxInt32
		for _, square := range squares {
			// if the square number is greater than the current number, break
			if i < square {
				break
			}
			// set the current number of squares to the minimum of the current number of
			currentMultiples := numbers[i]
			previousMultiples := numbers[i-square] + 1
			if currentMultiples < previousMultiples {
				numbers[i] = currentMultiples
			} else {
				numbers[i] = previousMultiples
			}
		}
	}
	return numbers[n]
}

func TestNumSquaresV2(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{
			n:        12,
			expected: 3,
		},
		{
			n:        13,
			expected: 2,
		},
	}
	for _, test := range tests {
		actual := NumSquaresV2(test.n)
		if actual != test.expected {
			t.Errorf("NumSquares(%d) = %d; expected %d", test.n, actual, test.expected)
		}
	}
}
