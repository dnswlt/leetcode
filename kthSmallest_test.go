package leetcode

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tree := ordered([]int{3, 5, 1, 4, 2})
	for i := 1; i <= 5; i++ {
		v := kthSmallest(tree, i)
		if v != i {
			t.Errorf("Expected %d, got %d\n", i, v)
		}
	}
}
