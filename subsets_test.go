package leetcode

import "testing"

func TestSubsets(t *testing.T) {
	s := subsets([]int{1, 2, 3})
	if len(s) != 78 {
		t.Error("Bad result", s)
	}
}
