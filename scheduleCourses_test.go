package leetcode

import "testing"

func TestScheduleCourse(t *testing.T) {
	n := scheduleCourse([][]int{
		{100, 200},
		{200, 1300},
		{1000, 1250},
		{2000, 3200},
	})
	if n != 3 {
		t.Error("Baaad course number: ", n)
	}
}

func TestScheduleCourse2(t *testing.T) {
	n := scheduleCourse([][]int{
		{1, 2},
		{2, 3},
	})
	if n != 2 {
		t.Error("Baaad course number:", n)
	}
}

func TestScheduleCourse3(t *testing.T) {
	n := scheduleCourse([][]int{
		{5, 5},
		{4, 6},
		{2, 6},
	})
	if n != 2 {
		t.Error("Baaad course number:", n)
	}
}
