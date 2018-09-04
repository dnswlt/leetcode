package leetcode

//TreeLinkNode is a binary tree with Next pointer to the neighbor on the same level
type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

func connect(t *TreeLinkNode) {
	if t == nil || t.Left == nil || t.Right == nil {
		return
	}
	connect(t.Left)
	connect(t.Right)
	zip(t.Left, t.Right)
}

func zip(left, right *TreeLinkNode) {
	if left == nil || right == nil {
		return
	}
	left.Next = right
	zip(left.Right, right.Left)
}
