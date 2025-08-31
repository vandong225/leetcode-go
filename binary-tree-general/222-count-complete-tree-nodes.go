package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 1

	count += countNodes(root.Left)
	if root.Right != nil {
		count += countNodes(root.Right)
	}

	return count
}
