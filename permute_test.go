package leetcode

import (
	"testing"
)

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}

func TestPermute(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ps := permute(v)
	if a, b := len(ps), fac(len(v)); a != b {
		t.Errorf("Expected %d items, got %d\n", b, a)
	}
}
