package matrix

import "math"

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i, r := range board {
		for j, c := range r {
			live := 0
			dead := 0
			for row := int(math.Max(0, float64(i-1))); row <= int(math.Min(float64(m-1), float64(i+1))); row++ {
				for col := int(math.Max(0, float64(j-1))); col <= int(math.Min(float64(n-1), float64(j+1))); col++ {
					if row == i && col == j {
						continue
					}
					bit := board[row][col] & 1
					if bit == 1 {
						live++
					} else {
						dead++
					}
				}
			}
			if (c == 1 && (live == 2 || live == 3)) || (c == 0 && live == 3) {
				board[i][j] |= 2
			}
		}
	}

	for i, r := range board {
		for j := range r {
			board[i][j] >>= 1
		}
	}
}
