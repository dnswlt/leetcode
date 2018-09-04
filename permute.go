package leetcode

func permute(nums []int) [][]int {
	var r [][]int
	if len(nums) == 0 {
		return r
	}
	perm(nums, 0, &r)
	return r
}

func perm(nums []int, sub int, buf *[][]int) {
	if sub == len(nums)-1 {
		v := make([]int, len(nums))
		copy(v, nums)
		*buf = append(*buf, v)
	}
	for i := sub; i < len(nums); i++ {
		nums[i], nums[sub] = nums[sub], nums[i]
		perm(nums, sub+1, buf)
		nums[i], nums[sub] = nums[sub], nums[i]
	}
}
