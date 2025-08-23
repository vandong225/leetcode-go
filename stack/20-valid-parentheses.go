package stack

func isValid(s string) bool {
	stack := make([]string, 0)

	for _, ch := range s {
		str := string(ch)
		if str == "(" || str == "{" || str == "[" {
			stack = append(stack, str)
		} else {
			if len(stack) == 0 {
				return false
			}
			n := len(stack) - 1
			topEl := stack[n]
			if !(str == ")" && topEl == "(" || str == "}" && topEl == "{" || str == "]" && topEl == "[") {
				return false
			}
			stack = stack[:n]
		}
	}

	return len(stack) == 0
}
