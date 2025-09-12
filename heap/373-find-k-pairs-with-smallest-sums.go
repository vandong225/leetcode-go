package heap

import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	res := [][]int{}

	if n == 0 || m == 0 || k == 0 {
		return res
	}

	h := &pairHeap{}
	heap.Init(h)

	for i := range nums1 {
		heap.Push(h, node{
			i: i, j: 0,
			sum: nums1[i] + nums2[0],
		})
	}

	for len(res) < k {
		first := heap.Pop(h).(node)
		a, b := nums1[first.i], nums2[first.j]

		res = append(res, []int{a, b})

		if first.j < m-1 {
			heap.Push(h, node{
				i:   first.i,
				j:   first.j + 1,
				sum: a + nums2[first.j+1],
			})
		}
	}

	return res
}

type node struct {
	i, j int
	sum  int
}

type pairHeap []node

func (h pairHeap) Len() int            { return len(h) }
func (h pairHeap) Less(a, b int) bool  { return h[a].sum < h[b].sum }
func (h pairHeap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(node)) }
func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
