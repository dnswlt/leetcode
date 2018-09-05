package leetcode

import (
	"fmt"
	"testing"
)

func iv(start, end int) Interval {
	return Interval{start, end}
}

func TestMergeIntervals(t *testing.T) {
	r := merge([]Interval{iv(1, 4), iv(0, 0)})
	if fmt.Sprint(r) != "[{0 0} {1 4}]" {
		t.Error("Not such a nice result", r)
	}
}

func TestMergeIntervals2(t *testing.T) {
	r := merge([]Interval{iv(1, 4), iv(2, 3)})
	if fmt.Sprint(r) != "[{1 4}]" {
		t.Error("Not such a nice result", r)
	}
}
