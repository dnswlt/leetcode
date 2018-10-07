package leetcode

import "testing"

func TestMaxProfitAssignment(t *testing.T) {
	p := maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7})
	if p != 100 {
		t.Error("Baaad profit: ", p)
	}
}
