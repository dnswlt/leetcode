package sorting

import (
	"container/heap"
	"math/rand"
)

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := rand.Intn(len(nums))
	m := partition(nums, pivot)
	qsort(nums[0:m])
	qsort(nums[m+1:])
}

func partition(nums []int, pivot int) int {
	if pivot > 0 {
		nums[pivot], nums[0] = nums[0], nums[pivot]
	}
	p := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= p {
			if i > l {
				nums[l], nums[i] = nums[i], nums[l]
			}
			l++
		}
	}
	if l > 1 {
		nums[0], nums[l-1] = nums[l-1], nums[0]
	}
	return l - 1
}

func msort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	m := len(nums) / 2
	msort(nums[:m])
	msort(nums[m:])
	merge(nums, m)
}

func merge(nums []int, m int) {
	q1 := make([]int, m)
	q2 := make([]int, len(nums)-m)
	copy(q1, nums)
	copy(q2, nums[m:])
	i := 0
	j := 0
	p := 0
	for i < len(q1) && j < len(q2) {
		if q1[i] < q2[j] {
			nums[p] = q1[i]
			i++
		} else {
			nums[p] = q2[j]
			j++
		}
		p++
	}
	for i < len(q1) {
		nums[p] = q1[i]
		i++
		p++
	}
	for j < len(q2) {
		nums[p] = q2[j]
		j++
		p++
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func hsort(nums []int) {
	cp := make([]int, len(nums))
	copy(cp, nums)
	h := IntHeap(cp)
	heap.Init(&h)
	i := 0
	for h.Len() > 0 {
		nums[i] = heap.Pop(&h).(int)
		i++
	}
}

func shellsort(nums []int) {
	gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}
	for _, gap := range gaps {
		for i := gap; i < len(nums); i++ {
			t := nums[i]
			j := i
			for ; j >= gap && nums[j-gap] > t; j -= gap {
				nums[j] = nums[j-gap]
			}
			nums[j] = t
		}
	}
}
