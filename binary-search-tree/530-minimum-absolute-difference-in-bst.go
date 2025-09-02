package binarysearchtree

import (
	"math"
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	nums := []int{}

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		n := len(nodes)
		for _, v := range nodes {
			nums = append(nums, v.Val)
			if v.Left != nil {
				nodes = append(nodes, v.Left)
			}
			if v.Right != nil {
				nodes = append(nodes, v.Right)
			}
		}

		nodes = nodes[n:]
	}

	slices.Sort(nums)
	cal := []int{}

	for i := len(nums) - 1; i > 0; i-- {
		cal = append(cal, int(math.Abs(float64(nums[i]-nums[i-1]))))
	}

	slices.Sort(cal)

	return cal[0]
}
