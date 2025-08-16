func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sInd := 0
	hasSame := false

	for i := 0; i < len(t); i++ {
		if s[sInd] == t[i] {
			sInd++
			hasSame = true
		}
		if sInd == len(s) {
			return hasSame
		}
	}

	return false
}