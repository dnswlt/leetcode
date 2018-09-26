package leetcode

import (
	"testing"
)

// combinationSums computes all possible combinations of items from nums (all positive) that add up to sum.
func TestCombinationSum(t *testing.T) {
	rg := make([]int, 40)
	for i := 0; i < len(rg); i++ {
		rg[i] = i + 1
	}
	sums := combinationSum(rg, 2*len(rg))
	if len(sums) != 10 {
		t.Error("Bad result", len(sums))
	}
}
