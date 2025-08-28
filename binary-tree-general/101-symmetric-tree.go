package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	return cmp(root.Left, root.Right)
}

func cmp(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	if r1.Val != r2.Val {
		return false
	}

	return cmp(r1.Left, r2.Right) && cmp(r1.Right, r2.Left)
}
