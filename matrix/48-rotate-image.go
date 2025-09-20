package matrix

func rotate(matrix [][]int) {
	n := len(matrix)

	for row := 0; row < n; row++ {
		for col := row + 1; col < n; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	for row := 0; row < n; row++ {
		for r, l := n-1, 0; l < r; l, r = l+1, r-1 {

		}
	}

}
