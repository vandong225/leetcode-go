package binarytreebfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	nv := []*TreeNode{root}

	for len(nv) != 0 {
		n := len(nv)
		cal := 0
		for i := 0; i < n; i++ {
			cal += nv[i].Val
			if nv[i].Left != nil {
				nv = append(nv, nv[i].Left)
			}
			if nv[i].Right != nil {
				nv = append(nv, nv[i].Right)
			}
		}

		nv = nv[n:]
		res = append(res, float64(cal)/float64(n))
	}

	return res
}
