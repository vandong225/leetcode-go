func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	max := 1
	c := 1

	for i, n := range nums {
		if i == 0 {
			continue
		}

		if n == nums[i-1] {
			continue
		}

		if n != nums[i-1]+1 {
			c = 1
		} else {
			c++

			if max < c {
				max = c
			}
		}

	}

	return max
}