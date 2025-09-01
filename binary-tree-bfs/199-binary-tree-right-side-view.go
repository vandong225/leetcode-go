package binarytreebfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	level := []*TreeNode{root}
	levels := [][]*TreeNode{level}

	for len(levels) != 0 {
		nodes := levels[0]

		res = append(res, nodes[len(nodes)-1].Val)

		newNodes := make([]*TreeNode, 0)
		for _, n := range nodes {
			if n.Left != nil {
				newNodes = append(newNodes, n.Left)
			}
			if n.Right != nil {
				newNodes = append(newNodes, n.Right)
			}
		}

		if len(newNodes) != 0 {
			levels = append(levels, newNodes)
		}

		levels = levels[1:]
	}

	return res

}
