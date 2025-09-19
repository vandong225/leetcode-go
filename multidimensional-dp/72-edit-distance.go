package multidimensionaldp

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)

	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= n2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				ins := dp[i-1][j] + 1
				del := dp[i][j-1] + 1
				rep := dp[i-1][j-1] + 1

				dp[i][j] = min3(ins, del, rep)
			}
		}
	}

	return dp[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a, b, c int) int { return min(a, min(b, c)) }
