package leetcode

// oddEvenList rearranges a list [1, 2, 3, 4, 5] to [1, 3, 5, 2, 4]. Odd indices come before even indices
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odds := head
	evens := head.Next
	evensHead := evens
	for {
		nextOdd := evens.Next
		if nextOdd == nil {
			// List ends at even
			odds.Next = evensHead
			break
		}
		if nextOdd.Next == nil {
			// List ends at odd
			evens.Next = nil
			odds.Next = nextOdd
			nextOdd.Next = evensHead
			break
		}
		odds.Next = nextOdd
		evens.Next = nextOdd.Next
		odds = odds.Next
		evens = evens.Next
	}
	return head
}
