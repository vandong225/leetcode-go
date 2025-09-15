package math

func plusOne(digits []int) []int {
	n := len(digits) - 1
	carry := 1
	for carry > 0 && n >= 0 {
		num := (digits[n] + carry)
		digits[n] = num % 10
		carry = num / 10
		n--
	}
	if digits[0] == 0 {
		res := make([]int, 0, len(digits)+1)
		res = append(res, 1)
		res = append(res, digits...)
		return res
	}

	return digits
}
