package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	n := firstMissingPositive([]int{1, 2, 3, 4})
	if n != 5 {
		t.Error("Not so great", n)
	}
}

func TestFirstMissingPositive2(t *testing.T) {
	n := firstMissingPositive([]int{1, 2, 2, 3})
	if n != 4 {
		t.Error("Not so great", n)
	}
}

func TestFirstMissingPositive3(t *testing.T) {
	n := firstMissingPositive([]int{3, 2, 1})
	if n != 4 {
		t.Error("Not so great", n)
	}
}

func TestFirstMissingPositive4(t *testing.T) {
	n := firstMissingPositive([]int{0})
	if n != 1 {
		t.Error("Not so great", n)
	}
}

func TestFirstMissingPositive5(t *testing.T) {
	n := firstMissingPositive([]int{1})
	if n != 2 {
		t.Error("Not so great", n)
	}
}

func TestFirstMissingPositive6(t *testing.T) {
	n := firstMissingPositive([]int{3})
	if n != 1 {
		t.Error("Not so great", n)
	}
}
