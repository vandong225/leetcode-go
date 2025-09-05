package graphgeneral

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, e := range prerequisites {
		a, b := e[0], e[1] // b -> a
		adj[b] = append(adj[b], a)
	}

	const (
		white = 0
		gray  = 1
		black = 2
	)
	color := make([]int, numCourses)

	var hasCycle func(u int) bool
	hasCycle = func(u int) bool {
		if color[u] == gray {
			return true
		} // back-edge
		if color[u] == black {
			return false
		} // done
		color[u] = gray
		for _, v := range adj[u] {
			if hasCycle(v) {
				return true
			}
		}
		color[u] = black
		return false
	}

	for i := 0; i < numCourses; i++ {
		if color[i] == white && hasCycle(i) {
			return false
		}
	}
	return true
}
