func minSubArrayLen(target int, nums []int) int {
	l, s := 0, 0
	ans := len(nums) + 1

	for r, num := range nums {
		s += num

		for s >= target {
			ans = min(ans, r-l+1)
			s -= nums[l]
			l++
		}
	}

	if ans == len(nums)+1 {
		return 0
	}

	return ans
}