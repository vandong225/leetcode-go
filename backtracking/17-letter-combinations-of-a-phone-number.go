package backtracking

var dict = map[rune][]rune{
	'2': []rune{'a', 'b', 'c'},
	'3': []rune{'d', 'e', 'f'},
	'4': []rune{'g', 'h', 'i'},
	'5': []rune{'j', 'k', 'l'},
	'6': []rune{'m', 'n', 'o'},
	'7': []rune{'p', 'q', 'r', 's'},
	'8': []rune{'t', 'u', 'v'},
	'9': []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	runes := []rune(digits)
	if len(runes) == 0 {
		return []string{}
	}

	res := []string{}

	digit := runes[0]

	for _, c := range dict[digit] {
		sub := letterCombinations(string(runes[1:]))
		str := string(c)
		if len(sub) == 0 {
			res = append(res, str)
		} else {

			for _, s := range sub {
				res = append(res, str+s)
			}
		}

	}

	return res
}
