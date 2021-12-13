package linkedlist

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func (n Node) GetVal() int {
	return n.Val
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(data int) {
	n := &Node{
		Val:  data,
		Next: l.head,
	}
	l.head = n
	l.length++
}

func (l *LinkedList) Append(data int) {
	n := &Node{
		Val: data,
	}
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for ; current.Next != nil; current = current.Next {
		}

		current.Next = n
	}
}

func (l LinkedList) Search(data int) bool {
	for current := l.head; current != nil; current = current.Next {
		if current.Val == data {
			return true
		}
	}
	return false
}

func (l *LinkedList) Delete(data int) error {
	if l.head.Val == data {
		l.head = l.head.Next
		l.length--
		return nil
	} else if l.head == nil {

	} else {
		current := l.head
		for ; current.Next != nil; current = current.Next {
			if current.Next.Val == data {
				current.Next = current.Next.Next
				l.length--
				return nil
			}
		}
	}
	return fmt.Errorf("tidak ditemukan")
}

func (l LinkedList) Print() {
	for current := l.head; current != nil; current = current.Next {
		fmt.Printf("%v ", current.Val)
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
	for ; current.Next != nil; current = current.Next {
	}
	return current.Val
}

func New() *LinkedList {
	return &LinkedList{}
}
