package queue

type CircularQueue struct {
	head  int
	tail  int
	queue []int
	size  int
}

func NewCircleQueue(k int) CircularQueue {
	return CircularQueue{
		head:  -1,
		tail:  -1,
		queue: make([]int, k),
		size:  k,
	}
}

func (cq *CircularQueue) EnQueue(value int) bool {
	if cq.IsFull() {
		// the queue is full
		return false
	}
	if cq.IsEmpty() {
		cq.head = 0
	}
	cq.tail = (cq.tail + 1) % cq.size
	cq.queue[cq.tail] = value
	return true
}

func (cq *CircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		// the queue is empty
		return false
	}
	//value := cq.queue[cq.tail]
	// don't need to reset the value as it will be overridden on enqueue
	//cq.queue[cq.head] = 0
	// check if the queue is empty and that was the last element
	if cq.head == cq.tail {
		cq.head = -1
		cq.tail = -1
		return true
	}
	cq.head = (cq.head + 1) % cq.size
	return true
}

func (cq *CircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.queue[cq.head]
}

func (cq *CircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.queue[cq.tail]
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.head == -1
}

func (cq *CircularQueue) IsFull() bool {
	//if cq.tail-cq.head == cq.size-1 {
	//	if cq.size == 1 {
	//		if cq.tail == -1 && cq.head == -1 {
	//			return false
	//		}
	//	}
	//	return true
	//}
	//if cq.tail-cq.head == -1 {
	//	return true
	//}
	//return false

	// much better!
	return ((cq.tail + 1) % cq.size) == cq.head
}
