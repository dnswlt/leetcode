package leetcode

import (
	"container/heap"
)

type freq struct {
	val  int
	freq int
}

type freqHeap []freq

func (h freqHeap) Len() int           { return len(h) }
func (h freqHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h freqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *freqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}
func (h *freqHeap) Push(f interface{}) {
	*h = append(*h, f.(freq))
}

func topKFrequent(nums []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	f := map[int]int{}
	for _, v := range nums {
		f[v] = f[v] + 1
	}
	h := &freqHeap{}
	for v, c := range f {
		heap.Push(h, freq{v, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	r := make([]int, k)
	for i := len(r) - 1; i >= 0; i-- {
		r[i] = heap.Pop(h).(freq).val
	}
	return r
}
