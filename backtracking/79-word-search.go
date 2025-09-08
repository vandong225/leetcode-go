package backtracking

func exist(board [][]byte, word string) bool {

	m := len(board)
	n := len(board[0])

	var dfs func(x, y, indStr int) bool

	dfs = func(x, y, indStr int) bool {
		if len(word) == indStr {
			return true
		}

		if x < 0 || y < 0 || x >= m || y >= n || word[indStr] != board[x][y] {
			return false
		}

		if len(word)-1 == indStr {
			return true
		}
		board[x][y] = '_'

		res := dfs(x, y+1, indStr+1) ||
			dfs(x+1, y, indStr+1) ||
			dfs(x, y-1, indStr+1) ||
			dfs(x-1, y, indStr+1)

		board[x][y] = word[indStr]
		return res
	}

	for x, r := range board {
		for y, _ := range r {
			res := dfs(x, y, 0)
			if res {
				return res
			}
		}
	}

	return false
}
