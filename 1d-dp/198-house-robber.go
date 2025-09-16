package ddp

import "math"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	maom := make([]int, n)

	maom[0] = nums[0]
	maom[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < n; i++ {
		maom[i] = int(math.Max(float64(nums[i]+maom[i-2]), float64(maom[i-1])))
	}

	return maom[n-1]
}
