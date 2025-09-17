package ddp

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 && dp[i-c] != math.MaxInt {
				dp[i] = int(math.Min(float64(dp[i-c]+1), float64(dp[i])))
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}
