func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, str := range strs {
		d := make([]int, 26)

		for _, ch := range str {
			d[ch-'a']++
		}

		var res strings.Builder

		for _, num := range d {
			res.WriteString(string(num))
		}

		dict[res.String()] = append(dict[res.String()], str)

	}

	anagrams := make([][]string, len(dict))

	i := 0
	for _, v := range dict {
		anagrams[i] = v
		i++
	}

	return anagrams
}