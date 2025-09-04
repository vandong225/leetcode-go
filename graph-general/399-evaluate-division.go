package graphgeneral

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	for i, r := range equations {
		x, y, v := r[0], r[1], values[i]
		if _, ok := graph[x]; !ok {
			graph[x] = make(map[string]float64)
		}
		if _, ok := graph[y]; !ok {
			graph[y] = make(map[string]float64)
		}
		graph[x][y] = v
		graph[y][x] = 1.0 / v
	}

	var dfs func(start, to string, visited map[string]bool) float64

	dfs = func(start, to string, visited map[string]bool) float64 {
		if v, ok := graph[start][to]; ok {
			return v
		}

		if _, ok := graph[start]; !ok {
			return -1
		}

		if start == to {
			return 1
		}
		visited[start] = true

		for k, v := range graph[start] {
			if _, ok := visited[k]; ok {
				continue
			}
			res := dfs(k, to, visited)
			if res != -1 {
				return res * v
			}
		}

		return -1
	}

	res := make([]float64, 0, len(queries))

	for _, q := range queries {
		from, to := q[0], q[1]
		dict := make(map[string]bool)
		res = append(res, dfs(from, to, dict))
	}

	return res
}
