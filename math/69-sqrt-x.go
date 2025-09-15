package math

func mySqrt(x int) int {
	l, h := 0, x
	ans := 0
	for l <= h {
		mid := (l + h) / 2

		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			l = mid + 1
			ans = mid
		} else {
			h = mid - 1
		}
	}

	return ans
}
