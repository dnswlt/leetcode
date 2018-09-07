package leetcode

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	lens := make([]int, l)
	for i := 0; i < l; i++ {
		lens[i] = 1
	}
	m := 1
	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && lens[i] <= lens[j] {
				lens[i] = lens[j] + 1
			}
		}
		if lens[i] > m {
			m = lens[i]
		}
	}
	return m
}
