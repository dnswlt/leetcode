package leetcode

func kthSmallest(root *TreeNode, k int) int {
	v, _ := kth(root, k)
	return v
}

func kth(root *TreeNode, k int) (v, n int) {
	if root == nil {
		return 0, k
	}
	v, l := kth(root.Left, k)
	if l == 0 {
		return v, l
	} else if l == 1 {
		return root.Val, 0
	}
	return kth(root.Right, l-1)
}
