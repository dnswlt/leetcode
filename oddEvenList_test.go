package leetcode

import "testing"

func TestOddEvenListEvenLength(t *testing.T) {
	l := fromSlice([]int{1, 2, 3, 4, 5, 6})
	r := toSlice(oddEvenList(l))
	if len(r) != 6 {
		t.Errorf("Length %d unexpected", len(r))
	}
	expected := []int{1, 3, 5, 2, 4, 6}
	for i := 0; i < len(r); i++ {
		if r[i] != expected[i] {
			t.Error("Baad result", r)
			return
		}
	}
}
func TestOddEvenListOddLength(t *testing.T) {
	l := fromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	r := toSlice(oddEvenList(l))
	if len(r) != 7 {
		t.Errorf("Length %d unexpected", len(r))
	}
	expected := []int{1, 3, 5, 7, 2, 4, 6}
	for i := 0; i < len(r); i++ {
		if r[i] != expected[i] {
			t.Error("Baad result", r)
			return
		}
	}
}
