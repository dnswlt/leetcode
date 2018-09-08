package leetcode

func minmult(h []int, i, j int) int {
	if h[i] < h[j] {
		return h[i] * (j - i)
	}
	return h[j] * (j - i)
}

func maxArea(height []int) int {
	l := len(height)
	left := 0
	right := l - 1
	mh := minmult(height, left, right)
outer:
	for right > left {
		v := minmult(height, left, right)
		if v > mh {
			mh = v
		}
		if height[left] <= height[right] {
			for i := left + 1; i < right; i++ {
				if height[i] > height[left] {
					left = i
					continue outer
				}
			}
			break
		} else {
			for i := right - 1; i > left; i-- {
				if height[i] > height[right] {
					right = i
					continue outer
				}
			}
			break
		}
	}
	return mh
}

// The naive solution. Always works, but in O(n**2)
func maxArea2(height []int) int {
	l := len(height)
	mh := minmult(height, 0, l-1)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			v := minmult(height, i, j)
			if v > mh {
				mh = v
			}
		}
	}
	return mh
}
