package leetcode

func sortColors(nums []int) {
	r, w, b := 0, 0, len(nums)-1
	for w <= b {
		switch nums[w] {
		case 0:
			nums[r], nums[w] = nums[w], nums[r]
			r++
			w++
		case 1:
			w++
		case 2:
			nums[w], nums[b] = nums[b], nums[w]
			b--
		}
	}
}
