func convert(s string, numRows int) string {

	if numRows <= 1 || numRows >= len(s) {
		return s
	}

	var res strings.Builder
	res.Grow(len(s))
	cycle := 2*numRows - 2

	for r := 0; r < numRows; r++ {
		for i := r; i < len(s); i += cycle {
			res.WriteByte(s[i])
			if r != 0 && r != numRows-1 {
				j := i + cycle - 2*r
				if j < len(s) {

					res.WriteByte(s[i+cycle-2*r])
				}
			}
		}
	}

	return res.String()
}
