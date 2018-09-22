package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	ls := toSlice(mergeKLists([]*ListNode{
		fromSlice([]int{1, 2, 6, 10}),
		fromSlice([]int{2, 3, 4, 7, 7, 11}),
		fromSlice([]int{}),
	}))
	if fmt.Sprint(ls) != "[1 2 2 3 4 6 7 7 10 11]" {
		t.Error("Baaad merge", ls)
	}
}
