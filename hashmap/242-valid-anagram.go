func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := make([]int, 26)
	for _, ch := range s {
		dict[ch-'a']++
	}

	for _, ch := range t {
		dict[ch-'a']--
		if dict[ch-'a'] < 0 {
			return false
		}

	}

	return true
}