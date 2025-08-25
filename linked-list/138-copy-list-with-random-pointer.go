package linkedlist

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var node Node
	p := &node

	dict := make(map[int]*Node)
	headDict := make(map[*Node]int)

	ind := 0

	for head != nil {
		p.Next = &Node{Val: head.Val, Next: nil, Random: head.Random}
		dict[ind] = p.Next
		headDict[head] = ind
		p = p.Next
		head = head.Next
		ind++
	}

	p = node.Next

	for p != nil {
		indRd, okRd := headDict[p.Random]
		if !okRd {
			p.Random = nil
		} else {
			node, _ := dict[indRd]
			p.Random = node
		}

		p = p.Next
	}

	return node.Next
}
