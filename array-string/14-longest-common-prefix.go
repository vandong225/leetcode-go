func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				if i == 0 {
					return ""
				}
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}


