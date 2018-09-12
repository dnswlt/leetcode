package leetcode

type index2d struct {
	i int
	j int
}

func floodFill(board [][]byte, i, j int, fill byte) {
	if board[i][j] == fill {
		return
	}
	r := len(board)
	c := len(board[0])
	v := board[i][j]
	q := []index2d{index2d{i, j}}
	for len(q) > 0 {
		x := q[len(q)-1]
		q = q[:len(q)-1]
		board[x.i][x.j] = fill
		if x.i > 0 && board[x.i-1][x.j] == v {
			q = append(q, index2d{x.i - 1, x.j})
		}
		if x.i < r-1 && board[x.i+1][x.j] == v {
			q = append(q, index2d{x.i + 1, x.j})
		}
		if x.j > 0 && board[x.i][x.j-1] == v {
			q = append(q, index2d{x.i, x.j - 1})
		}
		if x.j < c-1 && board[x.i][x.j+1] == v {
			q = append(q, index2d{x.i, x.j + 1})
		}
	}
}

func surroundedRegions(board [][]byte) {
	if len(board) < 3 || len(board[0]) < 3 {
		return
	}
	r := len(board)
	c := len(board[0])
	// O => $ for O touching boundary
	for i := 0; i < r; i++ {
		if board[i][0] == 'O' {
			floodFill(board, i, 0, '$')
		}
		if board[i][c-1] == 'O' {
			floodFill(board, i, c-1, '$')
		}
	}
	for j := 0; j < c; j++ {
		if board[0][j] == 'O' {
			floodFill(board, 0, j, '$')
		}
		if board[r-1][j] == 'O' {
			floodFill(board, r-1, j, '$')
		}
	}
	// Capture inner regions
	for i := 1; i < r-1; i++ {
		for j := 1; j < c-1; j++ {
			if board[i][j] == 'O' {
				floodFill(board, i, j, 'X')
			}
		}
	}
	// $ => O
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == '$' {
				board[i][j] = 'O'
			}
		}
	}
}

func solve(board [][]byte) {
	surroundedRegions(board)
}
