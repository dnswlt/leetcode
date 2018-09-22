package leetcode

import (
	"testing"
)

func rangeN(n int) []int {
	ns := make([]int, n)
	for i := 0; i < n; i++ {
		ns[i] = i
	}
	return ns
}

func validateCounts(t *testing.T, chooseFunc func([]int, int) []int) {
	counts := map[int]int{}
	for i := 0; i < 1000; i++ {
		nums := rangeN(100)
		ks := chooseFunc(nums, 10)
		for _, k := range ks {
			counts[k]++
		}
	}
	for i := 0; i < 100; i++ {
		if counts[i] == 0 {
			t.Errorf("No %d found", i)
		}
	}
}

func TestChooseK(t *testing.T) {
	validateCounts(t, chooseK)
}

func TestChooseKUnknownN(t *testing.T) {
	validateCounts(t, chooseKUnknownN)
}
