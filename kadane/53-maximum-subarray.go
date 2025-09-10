package kadane

func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cur+nums[i] > nums[i] { // cur+nums[i] làm tổng nhỏ đi -> cut bắt đầu lại từ nums[i]
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		if cur > best {
			best = cur
		}
	}
	return best
}
