package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	var r [][]int
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		ps := permute(nums[1:])
		for _, p := range ps {
			r = append(r, append(p, nums[0]))
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return r
}
