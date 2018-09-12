package leetcode

func moveNegativeLeft(nums []int) int {
	firstPositive := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[firstPositive], nums[i] = nums[i], nums[firstPositive]
			firstPositive++
		}
	}
	return firstPositive
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func firstMissingPositive(nums []int) int {
	start := moveNegativeLeft(nums)
	if start == len(nums) {
		return 1
	}
	pnums := nums[start:]
	for i := 0; i < len(pnums); i++ {
		j := abs(pnums[i]) - 1
		if j < len(pnums) && pnums[j] > 0 {
			pnums[j] = -pnums[j]
		}
	}
	for i := 0; i < len(pnums); i++ {
		if pnums[i] > 0 {
			return i + 1
		}
	}
	return len(pnums) + 1
}
