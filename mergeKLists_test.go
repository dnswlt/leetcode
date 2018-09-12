package leetcode

import "testing"

func TestMergeKLists(t *testing.T) {
	ls := toSlice(mergeKLists([]*ListNode{
		fromSlice([]int{1, 2, 6, 10}),
		fromSlice([]int{2, 3, 4, 7, 7, 11}),
		fromSlice([]int{}),
	}))
	t.Error(ls)
}
