package trie

func findWords(board [][]byte, words []string) []string {

	m, n := len(board), len(board[0])

	dictW := &WordSearchNode{}

	for _, w := range words {
		dictW.AddWord(w)
	}

	res := []string{}

	dict := make(map[string]bool)

	var dfs func(i, j int, node *WordSearchNode, visited [][]bool)

	dfs = func(i, j int, node *WordSearchNode, visited [][]bool) {
		if i >= m || j >= n || i < 0 || j < 0 || visited[i][j] {
			return
		}
		dummy := node.Next[board[i][j]-'a']

		if dummy == nil {
			return
		}

		_, ok := dict[dummy.Word]
		if dummy.Word != "" && !ok {
			res = append(res, dummy.Word)
			dict[dummy.Word] = true
		}

		visited[i][j] = true
		dfs(i+1, j, dummy, visited)
		dfs(i-1, j, dummy, visited)
		dfs(i, j+1, dummy, visited)
		dfs(i, j-1, dummy, visited)
		visited[i][j] = false
	}

	visited := make([][]bool, m)
	for k := 0; k < m; k++ {
		visited[k] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, dictW, visited)
		}
	}
	return res
}

type WordSearchNode struct {
	Word string
	Next [26]*WordSearchNode
}

func (node *WordSearchNode) AddWord(w string) {
	dummy := node

	for _, c := range w {
		if dummy.Next[c-'a'] == nil {
			dummy.Next[c-'a'] = &WordSearchNode{}
		}

		dummy = dummy.Next[c-'a']
	}
	dummy.Word = w
}
