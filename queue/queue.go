package queue

import "fmt"

const offside = -1

var (
	ErrQueueFull  = fmt.Errorf("queue is full")
	ErrQueueEmpty = fmt.Errorf("queue is empty")
)

type LinearQueue struct {
	front, rear int
	queue       []int
}

func NewLinear(size int) LinearQueue {
	return LinearQueue{
		front: offside,
		rear:  offside,
		queue: make([]int, size),
	}
}

func (lq *LinearQueue) Enqueue(addQueue int) error {
	if lq.IsFull() {
		return ErrQueueFull
	}

	if lq.front == offside {
		lq.front += 1
	}

	lq.rear += 1
	lq.queue[lq.rear] = addQueue
	return nil
}

func (lq *LinearQueue) Dequeue() error {
	if lq.IsEmpty() {
		return ErrQueueEmpty
	}

	lq.queue[lq.front] = 0
	if lq.front == lq.rear {
		lq.front, lq.rear = offside, offside
	} else {
		lq.front += 1
	}

	return nil
}

func (lq LinearQueue) Peek() (int, error) {
	if lq.IsEmpty() {
		return 0, ErrQueueEmpty
	}
	return lq.queue[lq.front], nil
}

func (lq LinearQueue) IsEmpty() bool {
	return lq.front == offside
}

func (lq LinearQueue) IsFull() bool {
	return lq.rear == len(lq.queue)-1
}
