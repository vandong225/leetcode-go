func lengthOfLongestSubstring(s string) int {
	best := 0
	l := 0
	chs := make(map[rune]int)

	for r, ch := range s {
		ind, ok := chs[ch]
		if ok && ind >= l {
			l = ind + 1
		}

		chs[ch] = r

		best = max(best, r-l+1)
	}

	return best
}