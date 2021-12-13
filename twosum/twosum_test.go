package twosum

import (
	"data_structur/linkedlist"
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{
		3, 4, 3,
	}
	res := TwoSum(nums, 6)
	if res[0] != 0 && res[1] != 2 {
		fmt.Println(res)
		t.FailNow()
	}
}

func TestLinkTwoSum(t *testing.T) {
	tesTable := []struct {
		l1 []int
		l2 []int
		ex []int
	}{
		{
			l1: []int{9, 9, 9, 9, 9, 9, 9},
			l2: []int{9, 9, 9, 9},
			ex: []int{8, 9, 9, 9, 0, 0, 0, 1},
		}, {
			l1: []int{3, 4, 9},
			l2: []int{4, 6, 9, 9, 9},
			ex: []int{7, 0, 9, 0, 0, 1},
		},
	}
	for _, v := range tesTable {
		link1 := listToLink(v.l1)
		link2 := listToLink(v.l2)
		res := LinkTwoSum(link1, link2)
		for _, expect := range v.ex {
			if expect != res.Val {
				t.FailNow()
			}
			res = res.Next
		}
	}
}

func listToLink(list []int) *linkedlist.Node {
	var reslNode, current *linkedlist.Node
	for _, v := range list {
		if reslNode == nil {
			reslNode = &linkedlist.Node{
				Val: v,
			}
			current = reslNode
		} else {
			current.Next = &linkedlist.Node{
				Val: v,
			}
			current = current.Next
		}
	}
	return reslNode
}
