package leetcode

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sgn := 1
	if dividend < 0 && divisor > 0 {
		sgn = -1
		dividend = -dividend
	} else if dividend > 0 && divisor < 0 {
		sgn = -1
		divisor = -divisor
	} else if dividend < 0 && divisor < 0 {
		dividend, divisor = -dividend, -divisor
	}
	if dividend < divisor {
		return 0
	}
	c := 0
	for dividend >= divisor {
		d, p := divisor, 1
		for d+d <= dividend {
			d = d + d
			p = p + p
		}
		dividend -= d
		c += p
	}
	if sgn < 0 {
		return -c
	}
	return c
}
