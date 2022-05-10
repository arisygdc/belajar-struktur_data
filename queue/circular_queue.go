package queue

type CircularQueue struct {
	front, rear int
	arr         []int
}

func NewCircular(queueLength int) CircularQueue {
	return CircularQueue{
		front: offside,
		rear:  offside,
		arr:   make([]int, queueLength),
	}
}

func (cq *CircularQueue) Enqueue(addQueue int) error {
	if cq.IsFull() {
		return ErrQueueFull
	}

	if cq.IsEmpty() {
		cq.front += 1
	}

	postion := cq.next(cq.rear)
	cq.arr[postion] = addQueue
	cq.rear = postion
	return nil
}

func (cq *CircularQueue) Dequeue() error {
	if cq.IsEmpty() {
		return ErrQueueEmpty
	}

	cq.arr[cq.front] = 0
	if cq.front == cq.rear {
		cq.front, cq.rear = offside, offside
		return nil
	}

	cq.front = cq.next(cq.front)
	return nil
}

func (cq CircularQueue) Peek() (int, error) {
	if cq.IsEmpty() {
		return 0, ErrQueueEmpty
	}
	return cq.arr[cq.front], nil
}

func (cq CircularQueue) IsEmpty() bool {
	return cq.front == offside
}

func (cq CircularQueue) IsFull() bool {
	return cq.next(cq.rear) == cq.front
}

func (cq CircularQueue) next(position int) int {
	return (position + 1) % len(cq.arr)
}
