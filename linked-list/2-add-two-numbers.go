package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res ListNode
	p := &res
	remain := 0
	for l1 != nil || l2 != nil || remain != 0 {
		vl1 := 0
		vl2 := 0
		if l1 != nil {
			vl1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			vl2 = l2.Val
			l2 = l2.Next
		}
		cal := vl1 + vl2 + remain
		v := cal % 10
		remain = cal / 10

		p.Next = &ListNode{Val: v, Next: nil}

		p = p.Next
	}

	return res.Next
}
