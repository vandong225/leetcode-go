package bitmanipulation

func singleNumber(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	res := 0
	for _, v := range nums {
		res ^= v // XOR
	}

	return res
}
