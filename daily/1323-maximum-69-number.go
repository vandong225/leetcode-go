func maximum69Number(num int) int {
	pow := 10
	for {
		if num/pow < 10 {
			break
		}
		pow *= 10
	}

	res := 0

	for num > 0 {

		d := num / pow
		if d == 6 {
			res += 9*pow + num - d*pow
			break
		}
		res += d * pow

		num -= d * pow
		pow = pow / 10
	}

	return res
}