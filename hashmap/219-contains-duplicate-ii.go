func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)

	for i, n := range nums {
		if j, ok := dict[n]; ok && int(math.Abs(float64(i-j))) < k {
			return true
		}

		dict[n] = i
	}

	return false
}