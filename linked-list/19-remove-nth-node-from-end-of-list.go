package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ind := 0
	dict := make(map[int]*ListNode)

	p := head
	for p != nil {
		dict[ind] = p
		p = p.Next
		ind++
	}
	ind--
	bfNode, bfNodeOk := dict[ind-n]
	afterNode, afterNodeOk := dict[ind-n+2]
	if !bfNodeOk {
		return afterNode
	}

	if !afterNodeOk {
		bfNode.Next = nil
	} else {
		bfNode.Next = afterNode
	}

	return head
}
