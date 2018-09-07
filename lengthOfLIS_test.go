package leetcode

import "testing"

func TestLengthOfLIS(t *testing.T) {
	l := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	if l != 4 {
		t.Error("Wrong", l)
	}
}
