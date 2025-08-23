package stack

import "math"

type MinStack struct {
	v    []int
	minS []int
}

func Constructor() MinStack {
	return MinStack{v: make([]int, 0), minS: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.v = append(this.v, val)
	min := val
	if len(this.minS) > 0 {
		min = int(math.Min(float64(val), float64(this.minS[len(this.minS)-1])))
	}
	this.minS = append(this.minS, min)
}

func (this *MinStack) Pop() {
	n := len(this.v)
	if n > 0 {
		this.v = this.v[:n-1]
		this.minS = this.minS[:n-1]
	}
}

func (this *MinStack) Top() int {
	n := len(this.v)
	if len(this.v) > 0 {
		return this.v[n-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	n := len(this.minS)
	if len(this.v) > 0 {
		return this.minS[n-1]
	}
	return 0
}
