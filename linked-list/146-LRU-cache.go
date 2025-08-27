package linkedlist

type LRUNode struct {
	K    int
	V    int
	Pre  *LRUNode
	Next *LRUNode
}

type LRUCache struct {
	Cap  int
	Dict map[int]*LRUNode
	Head *LRUNode
	Tail *LRUNode
}

func Constructor(capacity int) LRUCache {
	head := &LRUNode{}
	tail := &LRUNode{}

	head.Pre = tail
	tail.Next = head
	return LRUCache{Cap: capacity, Dict: make(map[int]*LRUNode, capacity), Head: head, Tail: tail}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Dict[key]; ok {
		this.MoveToHead(key)
		return node.V
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Dict[key]; ok {
		node.V = value
		this.MoveToHead(key)
		return
	}

	this.Dict[key] = &LRUNode{V: value, K: key}

	this.MoveToHead(key)

	if len(this.Dict) > this.Cap {
		this.EvictLRU()
	}
}

func (this *LRUCache) MoveToHead(key int) {
	node, ok := this.Dict[key]
	if !ok {
		return
	}

	if node.Next != nil {
		node.Next.Pre = node.Pre
	}

	if node.Pre != nil {
		node.Pre.Next = node.Next
	}

	node.Next = this.Head
	node.Pre = this.Head.Pre

	node.Next.Pre = node
	node.Pre.Next = node
}

func (this *LRUCache) EvictLRU() {

	lastNode := this.Tail.Next

	this.Tail.Next = lastNode.Next
	lastNode.Next.Pre = this.Tail

	delete(this.Dict, lastNode.K)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
