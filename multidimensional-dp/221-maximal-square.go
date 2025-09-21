package multidimensionaldp

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	max := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	min3 := func(a, b, c int) int { return min(a, min(b, c)) }

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				top, left, topLeft := dp[i-1][j], dp[i][j-1], dp[i-1][j-1]

				dp[i][j] = 1 + min3(top, left, topLeft)

				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max
}
