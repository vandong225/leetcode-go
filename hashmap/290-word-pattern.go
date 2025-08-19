func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}
	dict := make(map[string]string)
	dict2 := make(map[string]string)

	for i, ch := range pattern {
		eW, wOk := dict[string(ch)]
		eCh, chOk := dict2[words[i]]

		if wOk && eW != words[i] {
			return false
		}

		if chOk && eCh != string(ch) {
			return false
		}

		dict[string(ch)] = words[i]
		dict2[words[i]] = string(ch)
	}

	return true
}