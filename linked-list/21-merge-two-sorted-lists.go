package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var node ListNode

	p := &node

	for list1 != nil && list2 != nil {
		v := 0
		if list1.Val < list2.Val {
			v = list1.Val
			list1 = list1.Next

		} else {
			v = list2.Val
			list2 = list2.Next

		}
		p.Next = &ListNode{Val: v, Next: nil}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return node.Next
}
