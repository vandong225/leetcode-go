package backtracking

func generateParenthesis(n int) []string {
	res := []string{}

	var genParenthesis func(l, r int, parenthesis []rune)

	genParenthesis = func(l, r int, parenthesis []rune) {
		if len(parenthesis) == n*2 {
			res = append(res, string(parenthesis))
			return
		}

		if l < n {
			parenthesis = append(parenthesis, '(')
			genParenthesis(l+1, r, parenthesis)
			parenthesis = parenthesis[:len(parenthesis)-1]
		}

		if r < l {
			parenthesis = append(parenthesis, ')')
			genParenthesis(l, r+1, parenthesis)
			parenthesis = parenthesis[:len(parenthesis)-1]

		}

	}

	genParenthesis(0, 0, []rune{})

	return res

}
