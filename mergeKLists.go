package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		return &ListNode{l1.Val, merge2Lists(l1.Next, l2)}
	}
	return &ListNode{l2.Val, merge2Lists(l1, l2.Next)}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge2Lists(lists[0], mergeKLists(lists[1:]))
}
