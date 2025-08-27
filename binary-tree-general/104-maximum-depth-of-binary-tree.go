package binarytreegeneral

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(travel(root.Left, 1)), float64(travel(root.Right, 1))))
}

func travel(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	return int(math.Max(float64(travel(node.Left, depth+1)), float64(travel(node.Right, depth+1))))
}
