package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type state struct {
	t     *TreeNode
	depth int
}

func rev(r []int) {
	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	r := [][]int{}
	q := []state{state{root, 0}}
	for len(q) > 0 {
		hd := q[0]
		q = q[1:]
		if hd.t == nil {
			continue
		}
		if hd.depth < len(r) {
			r[hd.depth] = append(r[hd.depth], hd.t.Val)
		} else {
			r = append(r, []int{hd.t.Val})
		}
		d := hd.depth + 1
		q = append(q, state{hd.t.Left, d}, state{hd.t.Right, d})
	}
	for i := 1; i < len(r); i += 2 {
		rev(r[i])
	}
	return r
}
