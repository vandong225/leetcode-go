package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var newRNode *TreeNode
	if root.Left != nil {
		newRNode = invertTree(root.Left)
	}
	var newLNode *TreeNode

	if root.Right != nil {
		newLNode = invertTree(root.Right)
	}

	root.Right = newRNode
	root.Left = newLNode

	return root
}
