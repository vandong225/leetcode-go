package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// inorder: [left, root, right]
// numLeft = mid -1 - 0 +1 = mid
// numRight = n -1 - (mid +1) + 1 = n - mid -1
// postorder: [left, right, root]
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}

	root := postorder[n-1]

	mid := 0

	for i, v := range inorder {
		if v == root {
			mid = i
			break
		}
	}

	left := buildTree(inorder[:mid], postorder[:mid])
	right := buildTree(inorder[mid+1:], postorder[mid:n-1])

	return &TreeNode{Val: root, Left: left, Right: right}
}
