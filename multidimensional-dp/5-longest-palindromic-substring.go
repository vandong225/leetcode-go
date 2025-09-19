package multidimensionaldp

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)

	for i, _ := range s {
		dp[i] = make([]bool, n)
	}

	for i, _ := range s {
		dp[i][i] = true
	}

	maxi := 0
	mini := 0

	for i, c := range s {
		for j, sc := range s[:i] {
			if c != sc {
				continue
			}

			if j+1 <= i-1 && !dp[i-1][j+1] {
				continue
			}

			dp[i][j] = true

			if i-j > maxi-mini {
				maxi = i
				mini = j
			}
		}
	}

	return s[mini : maxi+1]
}
