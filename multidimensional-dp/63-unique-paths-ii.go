package multidimensionaldp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)

	dp := make([][]int, n)

	for i, v := range obstacleGrid {
		dp[i] = make([]int, len(v))
	}

	if obstacleGrid[0][0] == 0 {

		dp[0][0] = 1
	}

	for i, r := range obstacleGrid {
		for j, c := range r {
			if i == 0 && j == 0 {
				continue
			}
			if c == 1 {
				continue
			}
			topWay := 0
			leftWay := 0
			if i-1 >= 0 {
				leftWay = dp[i-1][j]
			}
			if j-1 >= 0 {
				topWay = dp[i][j-1]
			}

			dp[i][j] = topWay + leftWay
		}
	}

	return dp[n-1][len(dp[n-1])-1]
}
