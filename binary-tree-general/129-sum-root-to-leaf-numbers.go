package binarytreegeneral

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {

	return sumRoot2Leaf(root, 0)
}

func sumRoot2Leaf(root *TreeNode, sum int) int {
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	res := 0

	if root.Left != nil {
		res += sumRoot2Leaf(root.Left, sum)
	}
	if root.Right != nil {
		res += sumRoot2Leaf(root.Right, sum)
	}

	return res
}
