package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dict := make(map[int]bool)
	p := head

	for p != nil {
		if _, ok := dict[p.Val]; ok {
			p = p.Next
			continue
		}

		if p.Next != nil && p.Val == p.Next.Val {
			dict[p.Val] = true
		}

		p = p.Next
	}

	var node ListNode

	pNode := &node

	for head != nil {
		if _, ok := dict[head.Val]; ok {
			head = head.Next
			continue
		}

		pNode.Next = &ListNode{Val: head.Val, Next: nil}
		head = head.Next
		pNode = pNode.Next
	}

	return node.Next
}
