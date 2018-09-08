package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	m := maxArea([]int{3, 1, 7, 1, 1, 1, 1, 5, 3})
	if m != 25 {
		t.Error("Funny height", m)
	}
}
