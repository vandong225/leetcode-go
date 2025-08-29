package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := preorder[0]

	indRootInorder := 0

	for i, v := range inorder {
		if v == root {
			indRootInorder = i
			break
		}
	}

	left := buildTree(preorder[1:indRootInorder+1], inorder[:indRootInorder])
	right := buildTree(preorder[indRootInorder+1:], inorder[indRootInorder+1:])

	return &TreeNode{Val: root, Left: left, Right: right}
}
