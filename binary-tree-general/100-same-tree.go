package binarytreegeneral

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)

	if !left {
		return false
	}
	right := isSameTree(p.Right, q.Right)

	return right

}
