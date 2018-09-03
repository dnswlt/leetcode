package leetcode

import "testing"

func eq(m, n [][]int) bool {
	if len(m) != len(n) {
		return false
	}
	for i := 0; i < len(m); i++ {
		if len(m[i]) != len(n[i]) {
			return false
		}
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}
	return true
}

func TestSetZeroes1(t *testing.T) {
	m1 := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{1, 0, 1}}
	expected := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{0, 0, 0}}
	setZeroes(m1)
	if !eq(m1, expected) {
		t.Error("Matrix not as expected", m1)
	}
}

func TestSetZeroes2(t *testing.T) {
	m1 := [][]int{
		{0, 1, 1},
		{1, 1, 1},
		{1, 1, 1}}
	expected := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 1, 1}}
	setZeroes(m1)
	if !eq(m1, expected) {
		t.Error("Matrix not as expected", m1)
	}
}

func TestSetZeroes3(t *testing.T) {
	m1 := [][]int{
		{1},
		{0}}
	expected := [][]int{
		{0},
		{0}}
	setZeroes(m1)
	if !eq(m1, expected) {
		t.Error("Matrix not as expected", m1)
	}
}
