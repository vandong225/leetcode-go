package multidimensionaldp

func isInterleave(s1 string, s2 string, s3 string) bool {
	ns1 := len(s1)
	ns2 := len(s2)
	ns3 := len(s3)
	if ns1+ns2 != ns3 {
		return false
	}

	if ns1 == 0 {
		return s2 == s3
	}

	if ns2 == 0 {
		return s1 == s3
	}

	dp := make([][]bool, ns1+1)

	for i := 0; i <= ns1; i++ {
		dp[i] = make([]bool, ns2+1)
	}
	dp[0][0] = true

	for i := 1; i <= ns1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for i := 1; i <= ns2; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}

	for i := 1; i <= ns1; i++ {
		for j := 1; j <= ns2; j++ {
			down := dp[i-1][j] && s1[i-1] == s3[i+j-1]
			right := dp[i][j-1] && s2[j-1] == s3[i+j-1]

			dp[i][j] = down || right
		}
	}

	return dp[ns1][ns2]
}
