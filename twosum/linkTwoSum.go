package twosum

type ListNode struct {
	Val  int
	Next *ListNode
}

func LinkTwoSum(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode, tmp *ListNode
	var devider, val int
	val = l1.Val + l2.Val
	resultNode, devider = newNode(val, devider)
	current := resultNode

	for ; l1.Next != nil || l2.Next != nil; l1, l2 = l1.Next, l2.Next {
		if l1.Next == nil {
			l1.Next = &ListNode{
				Val: 0,
			}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{
				Val: 0,
			}
		}
		val = l1.Next.Val + l2.Next.Val
		tmp, devider = newNode(val, devider)
		current.Next = tmp
		current = current.Next
	}
	if devider > 0 {
		current.Next = &ListNode{
			Val: devider,
		}
	}
	return resultNode
}

func newNode(val int, devider int) (node *ListNode, divOut int) {
	val += devider
	if val >= 10 {
		node = &ListNode{
			Val: (val % 10),
		}
	} else {
		node = &ListNode{
			Val: val,
		}
	}
	divOut = val / 10
	return
}
