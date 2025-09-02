package binarysearchtree

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	nodes := []*TreeNode{root}
	nums := make([]int, 0, k)

	for len(nodes) > 0 {
		n := len(nodes)
		for _, v := range nodes {
			if len(nums) < k {
				nums = append(nums, v.Val)
				slices.Sort(nums)
			} else {
				if nums[k-1] > v.Val {
					nums[k-1] = v.Val
					slices.Sort(nums)
				}
			}
			if v.Left != nil {
				nodes = append(nodes, v.Left)
			}
			if v.Right != nil {
				nodes = append(nodes, v.Right)
			}
		}

		nodes = nodes[n:]
	}

	return nums[k-1]
}
