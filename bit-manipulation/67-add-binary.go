package bitmanipulation

import "strconv"

func addBinary(a string, b string) string {
	n, m := len(a)-1, len(b)-1
	res := ""
	carry := 0
	for n >= 0 || m >= 0 || carry > 0 {
		intA := 0
		intB := 0
		if n >= 0 {
			intA, _ = strconv.Atoi(string(a[n]))
		}

		if m >= 0 {
			intB, _ = strconv.Atoi(string(b[m]))
		}
		sum := intA + intB + carry
		res = strconv.Itoa(sum%2) + res
		carry = sum / 2

		n--
		m--
	}

	return res
}
