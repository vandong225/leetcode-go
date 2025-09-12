package heap1

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	h := MinHeap{}
	heap.Init(&h)

	for _, x := range nums {
		heap.Push(&h, x)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}
	return h[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
