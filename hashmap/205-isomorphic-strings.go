func isIsomorphic(s string, t string) bool {
	dict := make(map[rune]rune)
	dict2 := make(map[rune]rune)

	runeS := []rune(s)
	runeT := []rune(t)

	for i := 0; i < len(s); i++ {
		vs := dict[runeS[i]]
		vt := dict2[runeT[i]]
		if vs != rune(0) && vs != runeT[i] {
			return false
		}
		if vt != rune(0) && vt != runeS[i] {
			return false
		}

		dict[runeS[i]] = runeT[i]
		dict2[runeT[i]] = runeS[i]

	}

	return true
}