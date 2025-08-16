func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if i+j >= len(haystack) {
				break
			}
			if haystack[i+j] != needle[j] {
				break
			}

			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1

}