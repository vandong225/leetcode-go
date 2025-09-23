package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0

	dict := make(map[int]*ListNode)

	for head != nil {
		dict[n] = head
		n++
		head = head.Next
	}

	node := &ListNode{}
	dummy := node

	for i := 0; i < n; i += k {
		for j := i + k - 1; j >= i; j-- {
			if j >= n {
				dummy.Next = dict[i]
				break
			}
			dummy.Next = &ListNode{
				Val: dict[j].Val,
			}
			dummy = dummy.Next
		}

	}

	return node.Next
}
