package tree

import (
	"testing"
)

var root *TreeNode = &TreeNode{
	Val: 15,
	Left: &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 13,
			Left: &TreeNode{
				Val: 11,
			},
		},
	},
	Right: &TreeNode{
		Val: 21,
		Left: &TreeNode{
			Val: 19,
			Left: &TreeNode{
				Val: 17,
			},
		},
		Right: &TreeNode{
			Val: 23,
		},
	},
}

var root2 *TreeNode = &TreeNode{
	Val: 10,
	Left: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
		},
	},
	Right: &TreeNode{
		Val: 15,
		Left: &TreeNode{
			Val: 13,
		},
		Right: &TreeNode{
			Val: 18,
		},
	},
}

var root3 *TreeNode = &TreeNode{
	Val: 10,
	Left: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
		},
	},
	Right: &TreeNode{
		Val: 15,
		Right: &TreeNode{
			Val: 18,
		},
	},
}

var root4 *TreeNode = &TreeNode{
	Val: 18,
	Left: &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 12,
			},
		},
	},
	Right: &TreeNode{
		Val: 27,
		Left: &TreeNode{
			Val: 24,
			Left: &TreeNode{
				Val: 21,
			},
		},
		Right: &TreeNode{
			Val: 30,
		},
	},
}

func TestRangeSum(t *testing.T) {

	testTable := []struct {
		tree *TreeNode
		low  int
		high int
		excp int
	}{
		{tree: root, low: 19, high: 21, excp: 40},
		{tree: root2, low: 7, high: 15, excp: 32},
		{tree: root3, low: 6, high: 10, excp: 23},
		{tree: root4, low: 18, high: 24, excp: 63},
	}
	for i, v := range testTable {
		result := RangeSumBST(v.tree, v.low, v.high)
		t.Logf("%d. result %v expect %v\n", i+1, result, v.excp)
		if result != v.excp {
			t.Fail()
		}
	}
}
