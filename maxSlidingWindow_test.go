package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	ms := maxSlidingWindow([]int{1, 2, 3, 4, 5, 6}, 2)
	if fmt.Sprint(ms) != "[2 3 4 5 6]" {
		t.Error(ms)
	}
}

func TestMaxSlidingWindow2(t *testing.T) {
	ms := maxSlidingWindow([]int{5, 4, 3, 2, 1}, 1)
	if fmt.Sprint(ms) != "[5 4 3 2 1]" || true {
		t.Error(ms)
	}
}

func TestMaxSlidingWindow3(t *testing.T) {
	ms := maxSlidingWindow([]int{1, -1, 2, -2, 3, -3}, 2)
	if fmt.Sprint(ms) != "[1 2 2 3 3]" || true {
		t.Error(ms)
	}
}
