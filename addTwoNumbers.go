package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

func addTwo(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry > 0 {
			return &ListNode{carry, nil}
		}
		return nil
	}
	var v1, v2 int = 0, 0
	var n1, n2 *ListNode = nil, nil
	if l1 != nil {
		v1 = l1.Val
		n1 = l1.Next
	}
	if l2 != nil {
		v2 = l2.Val
		n2 = l2.Next
	}
	v := v1 + v2 + carry
	if v < 10 {
		return &ListNode{v, addTwo(n1, n2, 0)}
	}
	return &ListNode{v - 10, addTwo(n1, n2, 1)}

}
