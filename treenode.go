package leetcode

func insertSorted(val int, t *TreeNode) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	} else if val < t.Val {
		return &TreeNode{Val: t.Val, Left: insertSorted(val, t.Left), Right: t.Right}
	} else {
		return &TreeNode{Val: t.Val, Left: t.Left, Right: insertSorted(val, t.Right)}
	}
}

func ordered(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	return insertSorted(vals[0], ordered(vals[1:]))
}
