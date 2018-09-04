package leetcode

func subsets(nums []int) [][]int {
	l := uint(len(nums))
	n := uint(1) << l
	result := [][]int{}
	for i := uint(0); i < n; i++ {
		s := []int{}
		for b := uint(0); b < l; b++ {
			if i&(uint(1)<<b) != 0 {
				s = append(s, nums[b])
			}
		}
		result = append(result, s)
	}
	return result
}
