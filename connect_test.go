package leetcode

import (
	"fmt"
	"testing"
)

func mktree(depth int, val int) (t *TreeLinkNode, nextVal int) {
	if depth == 0 {
		return nil, val
	}
	left, nextVal := mktree(depth-1, val+1)
	right, nextVal := mktree(depth-1, nextVal)
	return &TreeLinkNode{Left: left, Right: right, Val: val}, nextVal
}

func printtree(t *TreeLinkNode) {
	if t == nil {
		fmt.Print(".")
		return
	}
	fmt.Print("(", t.Val, " ")
	printtree(t.Left)
	fmt.Print(" ")
	printtree(t.Right)
	fmt.Print(")")
}

func TestConnect(t *testing.T) {
	t1 := &TreeLinkNode{
		Val: 1,
		Left: &TreeLinkNode{
			Val: 2},
		Right: &TreeLinkNode{
			Val: 3}}
	connect(t1)
	if t1.Left.Next.Val != 3 {
		t.Error("Not right")
	}
}

func TestConnect2(t *testing.T) {
	t1, _ := mktree(4, 1)
	connect(t1)
	printtree(t1)
	if t1.Left.Next.Val != 9 {
		t.Error("Not right", t1.Left.Next.Val)
	}
	if t1.Left.Left.Next.Next.Next.Val != 13 {
		t.Error("Not right", t1.Left.Left.Next.Next.Next.Val)
	}
	if t1.Left.Left.Right.Next.Next.Next.Val != 11 {
		t.Error("Not right", t1.Left.Left.Right.Next.Next.Next.Val)
	}
}
