package stack

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			i, _ := strconv.Atoi(token)
			stack = append(stack, i)
			continue
		}

		cal := 0
		n := len(stack)
		a := stack[n-1]
		b := stack[n-2]
		stack = stack[:n-2]
		switch token {
		case "+":
			cal = a + b
		case "-":
			cal = b - a
		case "*":
			cal = a * b
		case "/":
			cal = b / a
		}

		stack = append(stack, cal)

	}

	return stack[0]
}
