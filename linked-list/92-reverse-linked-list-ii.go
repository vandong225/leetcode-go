package linkedlist

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dict := make(map[int]*ListNode)
	p := head
	ind := 1
	for ind <= right {
		dict[ind] = p
		ind++
		p = p.Next
	}

	for left < right {
		leftNode, _ := dict[left]
		rightNode, _ := dict[right]
		tempV := rightNode.Val
		rightNode.Val = leftNode.Val
		leftNode.Val = tempV

		left++
		right--
	}

	return head
}
