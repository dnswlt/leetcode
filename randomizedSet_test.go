package leetcode

import "testing"

func TestRandomizedSet(t *testing.T) {
	r := Constructor()
	b := r.Insert(1)
	if !b {
		t.Error("1 already there")
	}
	r.Insert(2)
	r.Insert(3)
	r.Remove(4)
	r.Remove(3)
	r.Remove(1)
	v := r.GetRandom()
	if v != 2 {
		t.Error("Funny random value", v)
	}
}
