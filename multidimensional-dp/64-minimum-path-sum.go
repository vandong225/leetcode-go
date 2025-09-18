package multidimensionaldp

import "math"

func minPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)

	for i, v := range grid {
		dp[i] = make([]int, len(v))
	}

	dp[0][0] = grid[0][0]

	for r := 0; r < n; r++ {
		for c := 0; c < len(grid[r]); c++ {
			if r == 0 && c == 0 {
				continue
			}
			v := 0
			if r-1 < 0 {
				v = dp[r][c-1]
			} else if c-1 < 0 {
				v = dp[r-1][c]
			} else {
				v = int(math.Min(float64(dp[r][c-1]), float64(dp[r-1][c])))
			}

			dp[r][c] = v + grid[r][c]
		}
	}

	return dp[n-1][len(dp[n-1])-1]
}
