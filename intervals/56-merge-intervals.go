package interval

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)

	start := 0
	for start < len(intervals) {

		min := intervals[start][0]
		max := intervals[start][1]

		for start < len(intervals) && max >= intervals[start][0] {
			min = int(math.Min(float64(min), float64(intervals[start][0])))
			max = int(math.Max(float64(max), float64(intervals[start][1])))
			start++
		}

		res = append(res, []int{min, max})

	}

	return res

}
