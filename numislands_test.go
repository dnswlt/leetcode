package leetcode

import "testing"

func TestNumIslands(t *testing.T) {
	n := numIslands([][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '0', '0', '1', '1'},
		[]byte{'0', '0', '0', '0', '1'},
		[]byte{'1', '1', '0', '0', '0'}})
	if n != 3 {
		t.Error("Other islands", n)
	}
}

func TestNumIslands2(t *testing.T) {
	n := numIslands([][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '1', '1', '1', '1'},
		[]byte{'0', '0', '0', '0', '1'},
		[]byte{'1', '1', '1', '1', '1'}})
	if n != 1 {
		t.Error("Other islands", n)
	}
}

func TestNumIslands3(t *testing.T) {
	n := numIslands([][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'}})
	if n != 1 {
		t.Error("Other islands", n)
	}
}
