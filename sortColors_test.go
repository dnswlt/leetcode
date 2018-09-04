package leetcode

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	if fmt.Sprint(colors) != "[0 0 1 1 2 2]" {
		t.Error("Not that shit", colors)
	}
}

func TestSortColors2(t *testing.T) {
	colors := []int{2, 0, 2, 1, 1, 0, 2}
	sortColors(colors)
	if fmt.Sprint(colors) != "[0 0 1 1 2 2 2]" {
		t.Error("Not that shit", colors)
	}
}

func TestSortColors3(t *testing.T) {
	colors := []int{2, 0, 2, 1, 1, 0, 2, 1}
	sortColors(colors)
	if fmt.Sprint(colors) != "[0 0 1 1 1 2 2 2]" {
		t.Error("Not that shit", colors)
	}
}
