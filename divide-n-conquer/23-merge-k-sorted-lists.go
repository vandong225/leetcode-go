package dividenconquer

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}

	node := &ListNode{}

	min := math.MaxInt
	minInd := -1
	for i := 0; i < k; i++ {
		if lists[i] != nil && lists[i].Val < min {
			min = lists[i].Val
			minInd = i
		}
	}

	if minInd == -1 {
		return nil
	}

	node.Val = min

	lists[minInd] = lists[minInd].Next

	node.Next = mergeKLists(lists)

	return node
}
