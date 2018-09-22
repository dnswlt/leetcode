package leetcode

import "math/rand"

// Not exactly leetcode, but an exercise from Skiena's book.

func chooseK(nums []int, k int) []int {
	// Choose K elements from nums at random, such that all items have same probability of being chosen.
	l := len(nums)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		j := rand.Intn(l - i)
		nums[i], nums[i+j] = nums[i+j], nums[i]
		result[i] = nums[i]
	}
	return result
}

func chooseKUnknownN(nums []int, k int) []int {
	vals := []int{}
	for i := 0; i < len(nums); i++ {
		// we pretent nums is a stream, so we only access "len(nums)" to iterate over this stream
		vals = append(vals, nums[i])
		j := rand.Intn(i + 1)
		if j > 0 {
			vals[i], vals[i-j] = vals[i-j], vals[i]
		}
	}
	return vals[:k]
}
