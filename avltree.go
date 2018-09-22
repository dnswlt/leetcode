package leetcode

import (
	"bufio"
	"fmt"
	"os"
)

// AVLTree represents an Adelson-Velsky-Landis tree
type AVLTree struct {
	root *avlNode
	size int
}

type avlNode struct {
	Val    int
	Height int
	Left   *avlNode
	Right  *avlNode
}

// Size returns the number of items in the AVLTree
func (t *AVLTree) Size() int {
	return t.size
}

// FindAVL finds the subtree whose root node value is v, or else nil.
func findAVL(t *avlNode, v int) *avlNode {
	if t == nil {
		return nil
	}
	if v == t.Val {
		return t
	}
	if v < t.Val {
		return findAVL(t.Left, v)
	}
	return findAVL(t.Right, v)
}

func height(t *avlNode) int {
	if t == nil {
		return 0
	}
	return t.Height
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (t *avlNode) rotl() {
	r := t.Right
	t.Val, r.Val = r.Val, t.Val
	t.Right = r.Right
	r.Left, r.Right = t.Left, r.Left
	t.Left = r
	r.fixHeight()
	t.fixHeight()
}

func (t *avlNode) fixHeight() {
	t.Height = max(height(t.Left), height(t.Right)) + 1
}

func (t *avlNode) rotr() {
	l := t.Left
	t.Val, l.Val = l.Val, t.Val
	t.Left = l.Left
	l.Left, l.Right = l.Right, t.Right
	t.Right = l
	l.fixHeight()
	t.fixHeight()
}

// Insert inserts element v into the tree and keeps the tree height-balanced
func (t *AVLTree) Insert(v int) {
	t.root = insertAVL(t.root, v)
	t.size++
}

// Values returns all values contained in the AVLTree, in order
func (t *AVLTree) Values() []int {
	vs := make([]int, t.size)
	avlValues(t.root, vs, 0)
	return vs
}

func avlValues(t *avlNode, vs []int, i int) int {
	if t == nil {
		return i
	}
	next := avlValues(t.Left, vs, i)
	vs[next] = t.Val
	return avlValues(t.Right, vs, next+1)
}

// WriteDot writes the tree as a .dot file to filename
func (t *AVLTree) WriteDot(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	defer w.Flush()
	w.WriteString("digraph {\n")
	q := []*avlNode{t.root}
	ids := map[*avlNode]int{t.root: 0}
	for len(q) > 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]
		fmt.Fprintf(w, "%d[label=\"%d\"]\n", ids[curr], curr.Val)
		if curr.Left != nil {
			ids[curr.Left] = len(ids)
			fmt.Fprintf(w, "%d->%d[label=\"L\"]\n", ids[curr], ids[curr.Left])
			q = append(q, curr.Left)
		}
		if curr.Right != nil {
			ids[curr.Right] = len(ids)
			fmt.Fprintf(w, "%d->%d[label=\"R\"]\n", ids[curr], ids[curr.Right])
			q = append(q, curr.Right)
		}
	}
	w.WriteString("}\n")
	return nil
}

func insertAVL(t *avlNode, v int) *avlNode {
	if t == nil {
		return &avlNode{v, 1, nil, nil}
	}
	if v <= t.Val {
		t.Left = insertAVL(t.Left, v)
		t.fixHeight()
		bal := height(t.Left) - height(t.Right)
		if bal <= 1 {
			return t
		}
		if height(t.Left.Left) < height(t.Left.Right) {
			t.Left.rotl()
		}
		// left-heavy tree
		t.rotr()
		return t
	}
	t.Right = insertAVL(t.Right, v)
	t.fixHeight()
	bal := height(t.Right) - height(t.Left)
	if bal <= 1 {
		return t
	}
	if height(t.Right.Right) < height(t.Right.Left) {
		t.Right.rotr()
	}
	t.rotl()
	return t
}
