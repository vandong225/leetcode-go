package intervals

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	x := points[0][1]

	c := 1

	for i := 1; i < len(points); i++ {
		if points[i][0] <= x {
			continue
		}
		x = points[i][1]
		c++
	}

	return c
}
