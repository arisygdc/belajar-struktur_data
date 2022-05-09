package queue

import (
	"log"
	"testing"
)

func TestLinearQueue(t *testing.T) {
	size := 5
	lq := NewLinear(size)
	for i := 1; i < 6; i++ {
		lq.Enqueue(i)
	}

	err := lq.Enqueue(10)
	if err == nil {
		log.Println("expected error")
		t.Fail()
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

	err = lq.Dequeue()
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
