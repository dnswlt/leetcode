package leetcode

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	i1, i2, m := 0, -1, 0
	for i := 1; i < len(nums); i++ {
		if i2 >= 0 && nums[i] > nums[i2] {
			return true
		} else if i2 == -1 && nums[i] > nums[i1] {
			i2 = i
		} else if i2 >= 0 && nums[i] > nums[i1] && nums[i] < nums[i2] {
			i2 = i
		}
		if nums[i] > nums[m] && (i2 < 0 || nums[i] < nums[i2]) {
			i1, i2 = m, i
		}
		if nums[i] < nums[m] {
			m = i
		}
	}
	return false
}
