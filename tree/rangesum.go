package tree

import "log"

type rangeSum struct {
	check map[int]int
}

func RangeSumBST(root *TreeNode, low int, high int) int {
	strs := &rangeSum{
		check: make(map[int]int),
	}
	result := strs.sumFinder(root, low, high, low) + strs.sumFinder(root, low, high, high)
	log.Println(strs.check)
	return result
}

func (strs *rangeSum) increment(val int, low int, high int) (result int) {
	if val >= low && val <= high && strs.check[val] == 0 {
		result += val
		strs.check[val] += 1
	}
	return
}

func (strs rangeSum) sumFinder(node *TreeNode, low int, high int, target int) (result int) {
	result += strs.increment(node.Val, low, high)
	if target < node.Val {
		if node.Left != nil {
			result += strs.sumFinder(node.Left, low, high, target)
		}
	} else if target > node.Val {
		if node.Right != nil {
			result += strs.sumFinder(node.Right, low, high, target)
		}
	}
	return
}
