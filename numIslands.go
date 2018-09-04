package leetcode

type unionfind struct {
	id   []int
	size []int
}

func ufMake(n int) unionfind {
	id := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		size[i] = 1
	}
	return unionfind{id, size}
}

func (uf unionfind) ufFind(i int) int {
	j := i
	for j != uf.id[j] {
		uf.id[j] = uf.id[uf.id[j]]
		j = uf.id[j]
	}
	return j
}

func (uf unionfind) ufUnion(p int, q int) {
	i := uf.ufFind(p)
	j := uf.ufFind(q)
	if uf.size[i] < uf.size[j] {
		uf.id[i] = j
		uf.size[j] += uf.size[i]
	} else {
		uf.id[j] = i
		uf.size[i] += uf.size[j]
	}
}

func (uf unionfind) ufPartitions() int {
	m := make(map[int]bool)
	for i := 0; i < len(uf.id); i++ {
		m[uf.ufFind(i)] = true
	}
	return len(m)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	uf := ufMake(len(grid) * len(grid[0]))
	r, c := len(grid), len(grid[0])
	numOnes := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				numOnes++
				if i > 0 && grid[i-1][j] == '1' {
					uf.ufUnion((i-1)*c+j, i*c+j)
				}
				if j > 0 && grid[i][j-1] == '1' {
					uf.ufUnion(i*c+j-1, i*c+j)
				}
			}
		}
	}
	return uf.ufPartitions() - (r*c - numOnes)
}
