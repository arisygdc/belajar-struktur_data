package linkedlist

import (
	"errors"
)

type DNode struct {
	prev *DNode
	next *DNode
	data interface{}
}

type DLinkedList struct {
	head *DNode
	tail *DNode
}

func NewD() *DLinkedList {
	return &DLinkedList{}
}

func (n DNode) Val() interface{} {
	return n.data
}

func (l *DLinkedList) AddToHead(data interface{}) {
	before := &DNode{
		data: data,
		next: nil,
		prev: nil,
	}

	if l.head == nil {
		l.head = before
		l.tail = before
	} else {
		before.next = l.head
		l.head.prev = before
		l.head = before
	}
}

func (l *DLinkedList) AddNode(data interface{}) {
	after := &DNode{
		data: data,
		next: nil,
		prev: nil,
	}

	l.tail = after
	if l.head == nil {
		l.head = after
	} else {
		curr := l.head
		for ; curr.next != nil; curr = curr.next {
		}
		curr.next = after
		after.prev = curr
	}
}

func (l DLinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *DLinkedList) DeleteAtHead() error {
	if l.IsEmpty() {
		return errors.New("cannot delete empty list")
	}

	l.head = l.head.next
	l.head.prev = nil
	return nil
}

func (l *DLinkedList) DeleteAtEnd() error {
	if l.IsEmpty() {
		return errors.New("cannot delete empty list")
	}

	l.tail = l.tail.prev
	l.tail.next = nil
	return nil
}

func (l *DLinkedList) Delete(target interface{}) error {
	head := l.head
	if l.IsEmpty() {
		return errors.New("cannot delete empty list")
	}
	curr := head
	if curr.Val() == target {
		l.DeleteAtHead()
		return nil
	}

	if l.tail.Val() == target {
		l.DeleteAtEnd()
		return nil
	}

	for curr = curr.next; curr.next != nil; curr = curr.next {
		if target == curr.Val() {
			curr.next.prev = curr.prev
			curr.prev.next = curr.next
			return nil
		}
	}

	return errors.New("target not found")
}

func (l DLinkedList) Array() []interface{} {
	var arr []interface{}
	for curr := l.head; curr != nil; curr = curr.next {
		arr = append(arr, curr.data)
	}

	return arr
}

func (l DLinkedList) ReverseArray() []interface{} {
	var arr []interface{}
	curr := l.tail
	for ; curr != nil; curr = curr.prev {
		arr = append(arr, curr.data)
	}

	return arr
}
