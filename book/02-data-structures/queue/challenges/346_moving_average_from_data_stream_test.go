package challenges

import (
	"fmt"
	"testing"
)

/*
Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.

Implement the MovingAverage class:

MovingAverage(int size) Initializes the object with the size of the window size.
double next(int val) Returns the moving average of the last size values of the stream.


Example 1:

Input
["MovingAverage", "next", "next", "next", "next"]
[[3], [1], [10], [3], [5]]
Output
[null, 1.0, 5.5, 4.66667, 6.0]

Explanation
MovingAverage movingAverage = new MovingAverage(3);
movingAverage.next(1); // return 1.0 = 1 / 1
movingAverage.next(10); // return 5.5 = (1 + 10) / 2
movingAverage.next(3); // return 4.66667 = (1 + 10 + 3) / 3
movingAverage.next(5); // return 6.0 = (10 + 3 + 5) / 3


Constraints:
- `1 <= size <= 1000`
- `-10^5 <= val <= 10^5`
- At most `10^4` calls will be made to next.
*/

type MovingAverage struct {
	queue []int
	size  int
	index int
}

func NewMovingAverage(size int) MovingAverage {
	return MovingAverage{
		queue: make([]int, 0),
		size:  size,
	}
}

func (ma *MovingAverage) Next(val int) float64 {
	// add the value to the queue
	if len(ma.queue) < ma.size {
		ma.queue = append(ma.queue, val)
	} else {
		ma.queue[ma.index] = val
	}
	// update the index
	ma.index = (ma.index + 1) % ma.size
	s := sum(ma.queue)
	l := float64(len(ma.queue))
	a := s / l
	//fmt.Println(ma.queue, s, l, a)
	return a
}

func sum[V int | float64](v []V) float64 {
	var s V
	for _, vi := range v {
		s += vi
	}
	return float64(s)
}

func TestMovingAverage(t *testing.T) {
	movingAverage := NewMovingAverage(3)
	o := movingAverage.Next(1) // return 1.0 = 1 / 1
	if o != 1 {
		t.Errorf("expected 1 got %f", o)
	}
	o = movingAverage.Next(10) // return 5.5 = (1 + 10) / 2
	if o != 5.5 {
		t.Errorf("expected 5.5 got %f", o)
	}
	o = movingAverage.Next(3) // return 4.66667 = (1 + 10 + 3) / 3
	if fmt.Sprintf("%f", o) != "4.666667" {
		t.Errorf("expected 4.666667 got %f", o)
	}
	o = movingAverage.Next(5) // return 6.0 = (10 + 3 + 5) / 3
	if o != 6 {
		t.Errorf("expected 6 got %f", o)
	}
}
