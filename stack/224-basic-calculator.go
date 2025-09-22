package stack

func calculate(s string) int {
	stack := []int{}
	num := 0
	sign := 1
	res := 0
	for i := 0; i < len(s); i++ {

		switch {
		case s[i] >= '0' && s[i] <= '9':
			num = num*10 + int(s[i]-'0')
		case s[i] == '+' || s[i] == '-':
			res += sign * num
			num = 0
			if s[i] == '-' {
				sign = -1
			} else {
				sign = 1
			}
		case s[i] == '(':
			stack = append(stack, res)
			stack = append(stack, sign)
			num = 0
			sign = 1
			res = 0
		case s[i] == ')':
			res += sign * num
			preSign := stack[len(stack)-1]
			preNum := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res = preNum + res*preSign
			num = 0
		case s[i] == ' ':
		}

	}
	res += sign * num
	return int(res)
}
