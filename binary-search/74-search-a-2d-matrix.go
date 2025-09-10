package binarysearch

func searchMatrix(matrix [][]int, target int) bool {

	left := 0
	right := len(matrix[0]) - 1
	for r := 0; r < len(matrix); r++ {
		if matrix[r][0] > target || matrix[r][len(matrix[0])-1] < target {
			continue
		}
		for left <= right {
			mid := (left + right) / 2
			v := matrix[r][mid]
			if v == target {
				return true
			}

			if v < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
