package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func TestAvlTree(t *testing.T) {
	var tree AVLTree
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	tree.WriteDot("avltree.dot")
}

func TestAvlTreeSeq(t *testing.T) {
	var tree AVLTree
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	vs := tree.Values()
	if len(vs) != 10 {
		t.Error("Funny length: ", len(vs))
	}
	if fmt.Sprint(vs) != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Error("Funny values: ", vs)
	}
}

func checkAllHeightsOK(t *avlNode) (int, bool) {
	if t == nil {
		return 0, true
	}
	hl, ok := checkAllHeightsOK(t.Left)
	if !ok {
		return -1, ok
	}
	hr, ok := checkAllHeightsOK(t.Right)
	if !ok {
		return -1, ok
	}
	h := max(hl, hr) + 1
	return h, h == t.Height
}

func checkBalance(t *avlNode) bool {
	if t == nil {
		return true
	}
	d := height(t.Left) - height(t.Right)
	return d >= -1 && d <= 1 && checkBalance(t.Left) && checkBalance(t.Right)
}

func TestHeight(t *testing.T) {
	const numElements = 100
	var tree AVLTree
	for i := 0; i < numElements/2; i++ {
		tree.Insert(i)
		tree.Insert(numElements - i)
	}
	h, ok := checkAllHeightsOK(tree.root)
	if !ok {
		t.Error("Errors in height found")
	}
	optHeight := int(math.Ceil(math.Log2(numElements)))
	if h > 2*optHeight {
		t.Errorf("Too high: %d, expected %d", h, optHeight)
	}
	if tree.Size() != numElements {
		t.Errorf("Unexpected size: %d", tree.Size())
	}
	if !checkBalance(tree.root) {
		t.Errorf("Tree not balanced")
	}
}
