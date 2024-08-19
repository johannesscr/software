package challenges

import "math"

type MinStack struct {
	len  int
	data []int
}

func Constructor() MinStack {
	ms := &MinStack{len: 0}
	return *ms
}

func (s *MinStack) Push(val int) {
	s.len++
	s.data = append(s.data, val)
}

func (s *MinStack) Pop() {
	s.data = s.data[:s.len-1]
	s.len--
}

func (s *MinStack) Top() int {
	return s.data[s.len-1]
}

func (s *MinStack) GetMin() int {
	minVal := math.MaxInt
	for _, v := range s.data {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}
