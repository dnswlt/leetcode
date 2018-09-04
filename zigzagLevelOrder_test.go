package leetcode

import (
	"fmt"
	"testing"
)

func TestZigZag1(t *testing.T) {
	tree := &TreeNode{Val: 1}
	r := zigzagLevelOrder(tree)
	if fmt.Sprint(r) != "[[1]]" {
		t.Error("Not so great", r)
	}
}

func TestZigZag2(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	r := zigzagLevelOrder(tree)
	if fmt.Sprint(r) != "[[1] [3 2]]" {
		t.Error("Not so great", r)
	}
}

func TestZigZag3(t *testing.T) {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	r := zigzagLevelOrder(tree)
	if fmt.Sprint(r) != "[[1] [3 2] [4 5]]" {
		t.Error("Not so great", r)
	}
}
