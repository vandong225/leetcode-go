package kadane

func maxSubarraySumCircular(nums []int) int {
	maxNormal, curMax, minSum, curMin := nums[0], nums[0], nums[0], nums[0]
	n := len(nums)
	total := nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		if curMax+x > x {
			curMax += x
		} else {
			curMax = x
		}
		if curMax > maxNormal {
			maxNormal = curMax
		}

		// Kadane min
		if curMin+x < x {
			curMin += x
		} else {
			curMin = x
		}
		if curMin < minSum {
			minSum = curMin
		}

		total += x
	}

	// all negative?
	if maxNormal < 0 {
		return maxNormal
	}

	// wrap case
	wrap := total - minSum
	if wrap > maxNormal {
		return wrap
	}
	return maxNormal
}
