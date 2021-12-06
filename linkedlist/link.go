package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func (n Node) GetVal() int {
	return n.data
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(data int) {
	n := &Node{
		data: data,
		next: l.head,
	}
	l.head = n
	l.length++
}

func (l *LinkedList) Append(data int) {
	n := &Node{
		data: data,
	}
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for ; current.next != nil; current = current.next {
		}

		current.next = n
	}
}

func (l LinkedList) Search(data int) bool {
	for current := l.head; current != nil; current = current.next {
		if current.data == data {
			return true
		}
	}
	return false
}

func (l *LinkedList) Delete(data int) error {
	if l.head.data == data {
		l.head = l.head.next
		l.length--
		return nil
	} else if l.head == nil {

	} else {
		current := l.head
		for ; current.next != nil; current = current.next {
			if current.next.data == data {
				current.next = current.next.next
				l.length--
				return nil
			}
		}
	}
	return fmt.Errorf("tidak ditemukan")
}

func (l LinkedList) Print() {
	for current := l.head; current != nil; current = current.next {
		fmt.Printf("%v ", current.data)
	}
	fmt.Println()
}

func (l LinkedList) GetLength() int {
	return l.length
}

func (l LinkedList) GetNode() *Node {
	return l.head
}

func (l LinkedList) GetTail() int {
	current := l.head
	for ; current.next != nil; current = current.next {
	}
	return current.data
}

func New() *LinkedList {
	return &LinkedList{}
}
