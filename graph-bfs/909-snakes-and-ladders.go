package graphbfs

import "container/list"

func snakesAndLadders(board [][]int) int {
	n := len(board)
	if n == 0 {
		return -1
	}
	target := n * n
	visited := make([]bool, target+1)
	q := list.New()
	q.PushBack(1)
	visited[1] = true

	move := 0

	for q.Len() > 0 {

		size := q.Len()

		for ; size > 0; size-- {
			curr := q.Remove(q.Front()).(int)
			if curr == target {
				return move
			}
			for roll := 1; roll <= 6 && curr+roll <= target; roll++ {
				row := n - 1 - (curr+roll-1)/n
				col := (curr + roll - 1) % n
				if (n-row)%2 == 0 {
					col = n - 1 - (curr+roll-1)%n
				}

				next := curr + roll

				if board[row][col] != -1 {
					next = board[row][col]
				}

				if !visited[next] {
					visited[next] = true
					q.PushBack(next)
				}

			}
		}

		move++
	}

	return -1
}
