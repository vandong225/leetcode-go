package graphgeneral

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	for _, sb := range prerequisites {
		a, b := sb[0], sb[1]
		indeg[a]++
		adj[b] = append(adj[b], a)
	}

	queue := []int{}
	res := make([]int, 0, numCourses)

	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		res = append(res, curr)
		for _, sb := range adj[curr] {
			indeg[sb]--
			if indeg[sb] == 0 {
				queue = append(queue, sb)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}

	return res
}
