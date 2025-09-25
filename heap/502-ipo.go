package heap

import "container/heap"

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	ch := &CapitalHeap{}
	ph := &ProfitHeap{}

	heap.Init(ch)
	heap.Init(ph)

	for i, c := range capital {
		heap.Push(ch, Capital{
			cap:    c,
			profit: profits[i],
		})
	}

	for k > 0 {
		for {
			if ch.Len() == 0 {
				break
			}
			cap := heap.Pop(ch).(Capital)
			if cap.cap > w {
				heap.Push(ch, cap)
				break
			}
			heap.Push(ph, cap.profit)
		}

		if ph.Len() == 0 {
			break
		}

		w += heap.Pop(ph).(int)

		k--
	}

	return w
}

type Capital struct {
	cap    int
	profit int
}

// IntMaxHeap implements heap.Interface and holds integers.
type CapitalHeap []Capital

func (h CapitalHeap) Len() int            { return len(h) }
func (h CapitalHeap) Less(a, b int) bool  { return h[a].cap < h[b].cap }
func (h CapitalHeap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *CapitalHeap) Push(x interface{}) { *h = append(*h, x.(Capital)) }
func (h *CapitalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type ProfitHeap []int

func (h ProfitHeap) Len() int            { return len(h) }
func (h ProfitHeap) Less(a, b int) bool  { return h[a] > h[b] }
func (h ProfitHeap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *ProfitHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *ProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
