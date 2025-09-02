package binarytreebfs

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	nodes := []*TreeNode{root}

	rl := true

	for len(nodes) > 0 {
		levelNodes := []int{}

		n := len(nodes)
		for i := 0; i < n; i++ {
			levelNodes = append(levelNodes, nodes[i].Val)

			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
		}

		nodes = nodes[len(levelNodes):]
		if rl {
			slices.Reverse(levelNodes)
		}
		res = append(res, levelNodes)
		rl = !rl
	}

	return res
}
