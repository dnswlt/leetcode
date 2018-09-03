package leetcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var l *ListNode
	l = addTwoNumbers(val(9), val(1))
	if v := of(l); v != 10 {
		t.Errorf("Expected %d, got %d", 10, v)
	}
	l = addTwoNumbers(val(999), val(1))
	if v := of(l); v != 1000 {
		t.Errorf("Expected %d, got %d", 1000, v)
	}
	l = addTwoNumbers(val(0), val(80))
	if v := of(l); v != 80 {
		t.Errorf("Expected %d, got %d", 80, v)
	}
}
