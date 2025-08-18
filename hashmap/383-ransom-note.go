func canConstruct(ransomNote string, magazine string) bool {
	c := make([]int, 26)
	for _, ch := range magazine {
		c[ch-'a']++
	}

	for _, ch := range ransomNote {
		v := c[ch-'a']
		if v <= 0 {
			return false
		}

		c[ch-'a'] = v - 1
	}

	return true
}