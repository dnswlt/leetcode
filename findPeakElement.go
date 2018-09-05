package leetcode

func findPeakElement(nums []int) int {
	return findPeakElem(nums, 0)
}

func findPeakElem(nums []int, offset int) int {
	if len(nums) == 1 {
		return offset
	}

	m := len(nums) / 2

	if (nums[m] > nums[m-1]) && (m == len(nums)-1 || nums[m] > nums[m+1]) {
		return offset + m
	}
	if nums[m] > nums[m-1] {
		return findPeakElem(nums[m+1:], offset+m+1)
	}
	return findPeakElem(nums[0:m], offset)
}
