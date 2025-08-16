func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1

	for {
		if start >= end {
			break
		}

		if !((s[start] >= 'a' && s[start] <= 'z') ||
			(s[start] >= '0' && s[start] <= '9')) {
			start++
			continue
		}

		if !((s[end] >= 'a' && s[end] <= 'z') ||
			(s[end] >= '0' && s[end] <= '9')) {
			end--
			continue
		}

		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}