package queue

import (
	"log"
	"testing"
)

func TestLinearQueue(t *testing.T) {
	size := 5
	lq := NewLinear(size)
	for i := 1; i < 6; i++ {
		if lq.IsFull() {
			t.Fail()
			log.Println("expected not full")
		}

		lq.Enqueue(i)

		if lq.IsEmpty() {
			t.Fail()
			log.Println("expected not empty")
		}
	}

	if lq.Enqueue(10) == nil {
		t.Fail()
		log.Println("expected full")
	}

	for i := 1; i < 6; i++ {
		peek, err := lq.Peek()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if peek != i {
			log.Printf("expected %d, got %d\n", i, peek)
			t.Fail()
		}

		lq.Dequeue()
	}

	err := lq.Dequeue()
	if err == nil {
		log.Println("expected error")
		t.Fail()
	}

	addToQueue := 100
	lq.Enqueue(addToQueue)
	peek, err := lq.Peek()
	if err != nil {
		log.Println("expected no error")
		t.Fail()
	}

	if addToQueue != peek {
		log.Printf("expected %d, got %d\n", addToQueue, peek)
		t.Fail()
	}
}

func TestCircularQueue(t *testing.T) {
	var lengTest = 6
	cQueue := NewCircular(lengTest)
	for i := 1; i <= lengTest; i++ {
		cQueue.Enqueue(i)
	}
	full := cQueue.IsFull()
	if !full {
		t.Fail()
		log.Println("expected queue is full")
	}

	for i := 1; i <= lengTest; i++ {
		peek, err := cQueue.Peek()
		if err != nil {
			t.Fail()
			log.Println(err)
		}

		if peek != i {
			t.Fail()
			log.Printf("expected %d, got %d\n", i, peek)
		}

		cQueue.Dequeue()
	}

	if !cQueue.IsEmpty() {
		t.Fail()
		log.Println("expected empty")
	}

	err := cQueue.Dequeue()
	if err == nil {
		t.Fail()
		log.Println("expected error: ", ErrQueueEmpty)
	}
}
