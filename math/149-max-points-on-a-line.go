package math

type Pair struct{ dy, dx int }

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return 1
	}
	return a
}

func normalize(dy, dx int) Pair {
	if dx == 0 { // vertical
		return Pair{1, 0}
	}
	if dy == 0 { // horizontal
		return Pair{0, 1}
	}
	g := gcd(dy, dx)
	dy /= g
	dx /= g
	if dx < 0 {
		dx = -dx
		dy = -dy
	}
	return Pair{dy, dx}
}

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	best := 0

	for i := 0; i < n; i++ {
		count := make(map[Pair]int, n)
		dup := 0
		localMax := 0

		for j := i + 1; j < n; j++ {
			dy := points[j][1] - points[i][1]
			dx := points[j][0] - points[i][0]
			if dy == 0 && dx == 0 {
				dup++
				continue
			}
			s := normalize(dy, dx)
			count[s]++
			if count[s] > localMax {
				localMax = count[s]
			}
		}
		cur := localMax + dup + 1
		if cur > best {
			best = cur
		}
	}
	return best
}
