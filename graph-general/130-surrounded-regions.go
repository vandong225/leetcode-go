package graphgeneral

func solve(board [][]byte) {
	var mark func(i, j int, from, to byte)

	mark = func(i, j int, from, to byte) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != from {
			return
		}

		board[i][j] = to
		mark(i+1, j, from, to)
		mark(i-1, j, from, to)
		mark(i, j+1, from, to)
		mark(i, j-1, from, to)
	}

	for i, r := range board {
		for j, _ := range r {
			if i == 0 || j == 0 || (i == len(board)-1) || (j == len(r)-1) {
				mark(i, j, 'O', '#')
			}
		}
	}

	for i, r := range board {
		for j, _ := range r {
			mark(i, j, 'O', 'X')
		}
	}

	for i, r := range board {
		for j, _ := range r {
			mark(i, j, '#', 'O')
		}
	}
}
