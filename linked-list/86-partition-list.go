package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	var newHead ListNode
	var tail ListNode

	ph := &newHead
	pt := &tail

	for head != nil {
		node := &ListNode{Val: head.Val, Next: nil}
		if head.Val < x {
			ph.Next = node

			ph = ph.Next
		} else {
			pt.Next = node
			pt = pt.Next
		}

		head = head.Next
	}

	ph.Next = tail.Next
	return newHead.Next
}
