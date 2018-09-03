package leetcode

func setZeroes(m [][]int) {
	r, c := len(m), len(m[0])
	col0 := false
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if m[i][j] == 0 {
				if j == 0 {
					col0 = true
				} else {
					m[0][j] = 0
				}
				m[i][0] = 0
			}
		}
	}
	// Zero rows
	for i := 1; i < r; i++ {
		if m[i][0] == 0 {
			for j := 1; j < c; j++ {
				m[i][j] = 0
			}
		}
	}
	// remember to zero out first row
	row0 := m[0][0] == 0
	// Zero columns
	for j := 0; j < c; j++ {
		if j == 0 && col0 || j > 0 && m[0][j] == 0 {
			for i := 0; i < r; i++ {
				m[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < c; j++ {
			m[0][j] = 0
		}
	}
}
