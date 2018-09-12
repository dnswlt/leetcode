package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	q := []int{}
	rs := []int{}
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	rs = append(rs, nums[q[0]])
	for i := k; i < len(nums); i++ {
		if q[0] <= i-k {
			q = q[1:]
		}
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		rs = append(rs, nums[q[0]])
	}
	return rs
}
