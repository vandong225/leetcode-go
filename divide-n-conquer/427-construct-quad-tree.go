package dividenconquer

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	n := len(grid)
	if n == 0 {
		return nil
	}
	node := &Node{}

	node.IsLeaf = isLeaf(grid)

	if !node.IsLeaf {

		topLeft := [][]int{}
		topRight := [][]int{}
		bottomLeft := [][]int{}
		bottomRight := [][]int{}

		for _, v := range grid[:n/2] {
			topLeft = append(topLeft, v[:n/2])
			topRight = append(topRight, v[n/2:])
		}

		for _, v := range grid[n/2:] {
			bottomLeft = append(bottomLeft, v[:n/2])
			bottomRight = append(bottomRight, v[n/2:])
		}

		node.TopLeft = construct(topLeft)
		node.TopRight = construct(topRight)
		node.BottomLeft = construct(bottomLeft)
		node.BottomRight = construct(bottomRight)

		node.Val = true
	} else {
		if grid[0][0] == 1 {
			node.Val = true
		}
	}

	return node

}

func isLeaf(grid [][]int) bool {
	if len(grid) <= 0 {
		return true
	}
	v := grid[0][0]
	for _, r := range grid {
		for _, c := range r {
			if c != v {
				return false
			}
		}
	}

	return true
}
