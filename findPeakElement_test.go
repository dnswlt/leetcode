package leetcode

import "testing"

func TestFindPeakElement(t *testing.T) {
	p := findPeakElement([]int{1, 1, 1, 1, 2, 3, 4, 3, 1, 1, 0})
	if p != 6 {
		t.Error("Not a peak", p)
	}
}

func TestFindPeakElement2(t *testing.T) {
	p := findPeakElement([]int{1, 2})
	if p != 1 {
		t.Error("Not a peak", p)
	}
}

func TestFindPeakElement3(t *testing.T) {
	p := findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
	if p != 5 {
		t.Error("Not a peak", p)
	}
}
