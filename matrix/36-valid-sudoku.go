package matrix

func isValidSudoku(board [][]byte) bool {
	rows := [9]int{}
	cols := [9]int{}
	boxes := [9]int{}

	for i, r := range board {
		for j, c := range r {
			if c == '.' {
				continue
			}
			bit := 1 << (c - 1)

			box := j/3 + (i/3)*3
			if rows[i]&bit != 0 || cols[j]&bit != 0 || boxes[box]&bit != 0 {
				return false
			}

			rows[i] |= bit
			cols[j] |= bit
			boxes[box] |= bit
		}
	}

	return true
}
