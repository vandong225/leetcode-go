package backtracking

func totalNQueens(n int) int {
	var backtracking func(row int, cols []bool, diags []bool, antiDiags []bool)
	res := 0
	backtracking = func(row int, cols []bool, diags []bool, antiDiags []bool) {
		if row == n {
			res++
			return
		}

		for c := 0; c < n; c++ {
			if cols[c] {
				continue
			}
			diag := row + c
			antiDiag := row - c + n - 1
			if diags[diag] || antiDiags[antiDiag] {
				continue
			}
			cols[c] = true
			diags[diag] = true
			antiDiags[antiDiag] = true
			backtracking(row+1, cols, diags, antiDiags)
			cols[c] = false
			diags[diag] = false
			antiDiags[antiDiag] = false
		}

	}
	cols := make([]bool, n)
	diags := make([]bool, 2*n-1)
	antiDiags := make([]bool, 2*n-1)

	backtracking(0, cols, diags, antiDiags)

	return res
}
