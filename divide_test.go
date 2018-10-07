package leetcode

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	args := [][]int{
		{12, 3, 4},
		{0, 1, 0},
		{-100, 9, -11},
		{math.MaxInt32, 1, math.MaxInt32},
		{math.MaxInt32 - 1, 1, math.MaxInt32 - 1},
		{-1, -1, 1},
		{math.MinInt32, -1, math.MaxInt32},
	}
	for _, a := range args {
		q := divide(a[0], a[1])
		if q != a[2] {
			t.Errorf("Bad quotient %d, expected %d\n", q, a[2])
		}
	}
}
