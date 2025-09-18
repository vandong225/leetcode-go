package multidimensionaldp

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i, v := range triangle {
		dp[i] = make([]int, len(v))
	}

	dp[0][0] = triangle[0][0]

	for r := 1; r < n; r++ {
		for c := 0; c < len(triangle[r]); c++ {
			v := 0
			if c-1 < 0 {
				v = dp[r-1][c]
			} else if c >= len(triangle[r-1]) {
				v = dp[r-1][c-1]
			} else {
				v = int(math.Min(float64(dp[r-1][c-1]), float64(dp[r-1][c])))
			}

			dp[r][c] = v + triangle[r][c]
		}
	}

	min := math.MaxInt

	for _, v := range dp[n-1] {
		if v < min {
			min = v
		}
	}
	return min
}
