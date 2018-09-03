package leetcode

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func val(n int) *ListNode {
	if n < 10 {
		return &ListNode{n, nil}
	}
	return &ListNode{n % 10, val(n / 10)}
}

func of(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val + 10*of(l.Next)
}

func toSlice(l *ListNode) []int {
	s := []int{}
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return s
}

func fromSlice(s []int) *ListNode {
	var l *ListNode
	for i := len(s) - 1; i >= 0; i-- {
		l = &ListNode{s[i], l}
	}
	return l
}
