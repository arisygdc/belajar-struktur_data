package linkedlist

import "fmt"

type DNode struct {
	prev *DNode
	next *DNode
	data interface{}
}

func (n DNode) Next() *DNode {
	return n.next
}

func (n DNode) Prev() *DNode {
	return n.prev
}

func (n DNode) Val() interface{} {
	return n.data
}

type DLinkedList struct {
	head   *DNode
	tail   *DNode
	length int
}

func (l *DLinkedList) Append(data interface{}) {
	n := &DNode{
		data: data,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n

		// Swap Tail
		n.prev = l.tail
		l.tail = n
	}
	l.length++
}

func (l DLinkedList) Print() {
	current := l.head
	for ; current != nil; current = current.next {
		fmt.Printf("%v ", current.data)
	}

	fmt.Println()
}

func (l DLinkedList) PrintReverse() {
	current := l.tail
	for ; current != nil; current = current.prev {
		fmt.Printf("%v ", current.data)
	}

	fmt.Println()
}

func (l *DLinkedList) Delete(target interface{}) error {
	current := l.head
	currentT := l.tail
	if l.length < 1 {
		return fmt.Errorf("tidak ditemukan")
	}
	if current.data == target {
		l.head = l.head.next
		for currentT.prev.prev != nil {
			currentT = currentT.prev
		}
		currentT.prev = currentT.prev.prev
		l.length--
		return nil
	}
	if currentT.data == target {
		l.tail = l.tail.prev
		for current.next.next != nil {
			current = current.next
		}
		current.next = current.next.next
		l.length--
		return nil
	}
	change := 0
	for current.next != nil {
		if current.next.data == target {
			current.next = current.next.next
			change++
			if change > 1 {
				return nil
			}
		}
		if currentT.prev.data == target {
			currentT.prev = currentT.prev.prev
			change++
			if change > 1 {
				return nil
			}
		}
		current = current.next
		currentT = currentT.prev
	}
	return fmt.Errorf("tidak ditemukan")
}

func (l DLinkedList) Head() *DNode {
	return l.head
}

func (l DLinkedList) Tail() *DNode {
	return l.tail
}

func NewD() *DLinkedList {
	return &DLinkedList{}
}
