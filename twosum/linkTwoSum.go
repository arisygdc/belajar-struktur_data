package twosum

import (
	"data_structur/linkedlist"
)

func LinkTwoSum(l1 *linkedlist.Node, l2 *linkedlist.Node) *linkedlist.Node {
	var resultNode, tmp *linkedlist.Node
	var devider, val int
	val = l1.Val + l2.Val
	resultNode, devider = newNode(val, devider)
	current := resultNode

	for ; l1.Next != nil || l2.Next != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Next == nil {
			l1.Next = &linkedlist.Node{
				Val: 0,
			}
		}
		if l2.Next == nil {
			l2.Next = &linkedlist.Node{
				Val: 0,
			}
		}
		val = l1.Next.Val + l2.Next.Val
		tmp, devider = newNode(val, devider)
		current.Next = tmp
		current = current.Next
	}
	if devider > 0 {
		current.Next = &linkedlist.Node{
			Val: devider,
		}
	}
	return resultNode
}

func newNode(val int, devider int) (node *linkedlist.Node, divOut int) {
	val += devider
	if val >= 10 {
		node = &linkedlist.Node{
			Val: (val % 10),
		}
	} else {
		node = &linkedlist.Node{
			Val: val,
		}
	}
	divOut = val / 10
	return
}
