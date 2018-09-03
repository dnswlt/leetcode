package leetcode

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	if !increasingTriplet([]int{1, 2, 2, 2, 2, 1, 3}) {
		t.Error("Expected a triplet")
	}
	if !increasingTriplet([]int{1, 2, 3}) {
		t.Error("Expected a triplet")
	}
	if increasingTriplet([]int{1, 2, 2, 2, 2, 1}) {
		t.Error("Expected no triplet")
	}
	if !increasingTriplet([]int{10, 9, 8, 1, 10, 9, 8, 2, 10, 3}) {
		t.Error("Expected a triplet")
	}
	if increasingTriplet([]int{10, 9, 8, 1, 7, 6, 5, 2}) {
		t.Error("Expected a triplet")
	}
	if !increasingTriplet([]int{1, 2, -1, 3}) {
		t.Error("Expected a triplet")
	}
}
