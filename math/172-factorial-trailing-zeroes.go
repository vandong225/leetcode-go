package math

func trailingZeroes(n int) int {
	cnt5 := 0
	for n > 0 {
		n /= 5
		cnt5 += n
	}

	return cnt5
}
