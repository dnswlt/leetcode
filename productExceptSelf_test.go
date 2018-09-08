package leetcode

import "testing"

func TestProductExceptSelf(t *testing.T) {
	r := productExceptSelf([]int{1, 2, 3, 4, 5})
	expected := []int{120, 60, 40, 30, 24}
	for i := 0; i < len(expected); i++ {
		if r[i] != expected[i] {
			t.Error("Unexpected", r, expected)
		}
	}
}

func TestProductExceptSelf2(t *testing.T) {
	r := productExceptSelf([]int{2, 2, 2, 2, 2})
	for i := 0; i < len(r); i++ {
		if r[i] != 16 {
			t.Error("Unexpected", r, 16)
		}
	}
}
