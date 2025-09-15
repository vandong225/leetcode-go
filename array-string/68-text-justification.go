package arraystring

import "strings"

func fullJustify(words []string, maxWidth int) []string {

	tempWords := []string{}
	lenWords := 0

	arrWords := [][]string{}
	res := []string{}
	for _, w := range words {
		if len(tempWords)+lenWords+len(w) > maxWidth {
			arrWords = append(arrWords, tempWords)
			tempWords = []string{w}
			lenWords = len(w)
		} else {
			tempWords = append(tempWords, w)
			lenWords += len(w)
		}
	}
	arrWords = append(arrWords, tempWords)

	for j, w := range arrWords {
		if j != len(arrWords)-1 {

			lenW := 0
			n := len(w)
			i := 0
			for lenW < maxWidth {
				if i < n {
					lenW += len(w[i%n])
				}
				if (i%n != n-1 || n == 1) && lenW < maxWidth {
					w[i%n] = w[i%n] + " "
					lenW++
				}
				i++
			}

			res = append(res, strings.Join(w, ""))
		} else {
			str := strings.Join(w, " ")
			str += strings.Repeat(" ", maxWidth-len(str))
			res = append(res, str)
		}
	}

	return res
}
