func isHappy(n int) bool {

	strN := strconv.Itoa(n)
	dict := make(map[int]bool)

	for {
		sum := 0
		for _, num := range strN {
			sum += int(math.Pow(float64(num-'0'), 2))
		}

		strN = strconv.Itoa(sum)

		if sum == 1 {
			return true
		}

		if _, ok := dict[sum]; ok {
			return false
		}

		dict[sum] = true
	}

	return n == 1

}