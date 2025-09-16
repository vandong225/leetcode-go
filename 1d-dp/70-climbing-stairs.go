package ddp

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	ways := make([]int, n)
	ways[0] = 1
	ways[1] = 2

	for i := 2; i < n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n-1]
}
