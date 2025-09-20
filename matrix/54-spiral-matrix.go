package matrix

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0, m*n)

	top, left, right, bottom := 0, 0, n-1, m-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {

			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {

			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res

}
