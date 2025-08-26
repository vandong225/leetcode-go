package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	sz := 0

	p := head

	dict := make(map[int]*ListNode)

	for p != nil {
		dict[sz] = p
		p = p.Next
		sz++
	}

	var node ListNode

	rotate := k % sz

	res := &node

	for i := 0; i < sz; i++ {
		el, _ := dict[(sz-rotate+i)%sz]
		res.Next = &ListNode{Val: el.Val, Next: nil}
		res = res.Next
	}

	return node.Next
}
