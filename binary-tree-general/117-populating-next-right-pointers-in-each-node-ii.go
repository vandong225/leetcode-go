package binarytreegeneral

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	level := root

	for level != nil {
		dummy := &Node{}
		tail := dummy

		for curr := level; curr != nil; curr = curr.Next {
			if curr.Left != nil {
				tail.Next = curr.Left
				tail = tail.Next
			}

			if curr.Right != nil {
				tail.Next = curr.Right
				tail = tail.Next
			}
		}

		level = dummy.Next
	}

	return root
}
