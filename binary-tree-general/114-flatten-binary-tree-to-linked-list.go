package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		lastRightOfLeft := root.Left

		for lastRightOfLeft.Right != nil {
			lastRightOfLeft = lastRightOfLeft.Right
		}

		lastRightOfLeft.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	flatten(root.Right)
}
