package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func minIndex(lists []*ListNode) int {
	min := -1
	for j := 0; j < len(lists); j++ {
		if lists[j] != nil {
			if min == -1 || lists[j].Val < lists[min].Val {
				min = j
			}
		}
	}
	return min
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	i := minIndex(lists)
	if i == -1 {
		return nil
	}
	v := lists[i].Val
	lists[i] = lists[i].Next
	return &ListNode{v, mergeKLists(lists)}
}
