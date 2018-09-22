package sorting

import (
	"math/rand"
	"sort"
	"testing"
)

func TestIsSorted(t *testing.T) {
	b := isSorted([]int{})
	if !b {
		t.Error("Empty slice is sorted")
	}
	b = isSorted([]int{1})
	if !b {
		t.Error("Singleton slice is sorted")
	}
	b = isSorted([]int{1, 2, 3, 4})
	if !b {
		t.Error("Ascending slice is sorted")
	}
	b = isSorted([]int{1, 1, 1, 1})
	if !b {
		t.Error("Constant slice is sorted")
	}
	b = isSorted([]int{1, 2, 3, 2})
	if b {
		t.Error("Not sorted slice is not sorted")
	}
}

func testSortAlgo(t *testing.T, nums []int, sorter func([]int)) {
	copy1 := make([]int, len(nums))
	copy(copy1, nums)
	sorter(copy1)
	copy2 := make([]int, len(nums))
	copy(copy2, nums)
	sort.Ints(copy2)
	for i := 0; i < len(nums); i++ {
		if copy1[i] != copy2[i] {
			t.Errorf("Slice not properly sorted. Got %v, expected %v\n", copy1, copy2)
		}
	}
}

func TestQsort(t *testing.T) {
	testSortAlgo(t, []int{1, 2, 3}, qsort)
	testSortAlgo(t, []int{1, 2, 3, 3}, qsort)
	testSortAlgo(t, []int{3, 2, 1}, qsort)
	testSortAlgo(t, []int{1, 1, 1}, qsort)
	testSortAlgo(t, []int{1}, qsort)
}

func TestMsort(t *testing.T) {
	testSortAlgo(t, []int{1, 2, 3}, msort)
	testSortAlgo(t, []int{1, 2, 3, 3}, msort)
	testSortAlgo(t, []int{3, 2, 1}, msort)
	testSortAlgo(t, []int{1, 1, 1}, msort)
	testSortAlgo(t, []int{1}, msort)
}

func TestHsort(t *testing.T) {
	testSortAlgo(t, []int{1, 2, 3}, hsort)
	testSortAlgo(t, []int{1, 2, 3, 3}, hsort)
	testSortAlgo(t, []int{3, 2, 1}, hsort)
	testSortAlgo(t, []int{1, 1, 1}, hsort)
	testSortAlgo(t, []int{1}, hsort)
}

func TestShellsort(t *testing.T) {
	testSortAlgo(t, []int{1, 2, 3}, shellsort)
	testSortAlgo(t, []int{1, 2, 3, 3}, shellsort)
	testSortAlgo(t, []int{3, 2, 1}, shellsort)
	testSortAlgo(t, []int{1, 1, 1}, shellsort)
	testSortAlgo(t, []int{1}, shellsort)
}

func randSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Int()
	}
	return s
}

func BenchmarkQsort(b *testing.B) {
	s := randSlice(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := make([]int, len(s))
		copy(t, s)
		qsort(s)
	}
}

func BenchmarkMsort(b *testing.B) {
	s := randSlice(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := make([]int, len(s))
		copy(t, s)
		msort(s)
	}
}

func BenchmarkHsort(b *testing.B) {
	s := randSlice(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := make([]int, len(s))
		copy(t, s)
		hsort(s)
	}
}

func BenchmarkShellsort(b *testing.B) {
	s := randSlice(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := make([]int, len(s))
		copy(t, s)
		shellsort(s)
	}
}
