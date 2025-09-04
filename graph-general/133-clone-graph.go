package graphgeneral

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	queue := []*Node{node}
	visited := make(map[int]*Node)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if _, ok := visited[curr.Val]; ok {
			continue
		}

		newNode := &Node{Val: curr.Val, Neighbors: make([]*Node, 0, len(curr.Neighbors))}
		for _, n := range curr.Neighbors {
			if ne, ok := visited[n.Val]; ok {
				ne.Neighbors = append(ne.Neighbors, newNode)
				newNode.Neighbors = append(newNode.Neighbors, ne)
			} else {
				queue = append(queue, n)
			}
		}

		visited[curr.Val] = newNode

	}

	res := visited[1]

	return res

}
