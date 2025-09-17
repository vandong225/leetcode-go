package ddp

import "math"

func lengthOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)

	max := 1

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))

				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}

	return max
}
