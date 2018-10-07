package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	m := maxProfit([]int{1, 2, 3, 0, 2})
	if m != 3 {
		t.Error("Baad result: ", m)
	}
}
