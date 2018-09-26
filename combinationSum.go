package leetcode

import "sort"

// combinationSums computes all possible combinations of items from nums (all positive) that add up to sum.
func combinationSum(nums []int, sum int) [][]int {
	sort.Ints(nums)
	if len(nums) == 0 || nums[0] > sum {
		return [][]int{}
	}
	r := [][]int{}
	buf := []int{}
	combinationSumBT(nums, sum, 0, buf, &r)
	return r
}

func combinationSumBT(nums []int, sum int, i int, buf []int, r *[][]int) {
	if sum == 0 {
		sol := make([]int, len(buf))
		copy(sol, buf)
		*r = append(*r, sol)
		return
	}
	for j := i; j < len(nums); j++ {
		if nums[j] <= sum {
			combinationSumBT(nums, sum-nums[j], j, append(buf, nums[j]), r)
		}
	}
}
