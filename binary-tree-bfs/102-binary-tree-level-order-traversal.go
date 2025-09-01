package binarytreebfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	nodes := []*TreeNode{root}

	for len(nodes) != 0 {
		n := len(nodes)
		vl := []int{}
		for i := 0; i < n; i++ {
			vl = append(vl, nodes[i].Val)

			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}

			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}

		nodes = nodes[n:]
		res = append(res, vl)

	}

	return res
}
