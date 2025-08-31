package binarytreegeneral

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// [mid, left, right]
// inorder: [left mid right]
type BSTIterator struct {
	Arr   []*TreeNode
	Index int
}

func Constructor(root *TreeNode) BSTIterator {

	return BSTIterator{Arr: flat(root), Index: -1}
}

func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return math.MinInt
	}
	this.Index++
	return this.Arr[this.Index].Val
}

func (this *BSTIterator) HasNext() bool {
	return this.Index < len(this.Arr)-1
}

func flat(root *TreeNode) []*TreeNode {
	arr := make([]*TreeNode, 0)
	if root == nil {
		return arr
	}

	arr = append(arr, flat(root.Left)...)
	arr = append(arr, root)

	arr = append(arr, flat(root.Right)...)
	return arr
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
