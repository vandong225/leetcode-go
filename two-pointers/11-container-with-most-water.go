func maxArea(height []int) int {
	maxS := 0
	start := 0
	end := len(height) - 1

	for start < end {
		newS := min(height[start], height[end]) * (end - start)
		if newS > maxS {
			maxS = newS
		}

	}

	return maxS
}
