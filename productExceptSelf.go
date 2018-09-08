package leetcode

func productExceptSelf(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	prod := make([]int, l)
	p := 1
	for i := 0; i < l; i++ {
		prod[i] = p
		p *= nums[i]
	}
	p = 1
	for i := l - 1; i >= 0; i-- {
		prod[i] *= p
		p *= nums[i]
	}
	return prod
}
