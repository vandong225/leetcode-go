package backtracking

func combine(n int, k int) [][]int {
	res := [][]int{}
	var backtracking func(start, remain int, currComb []int)

	backtracking = func(start, remain int, currComb []int) {
		if remain == 0 {
			dst := make([]int, k)
			copy(dst, currComb)
			res = append(res, dst)
			return
		}

		for i := start; i <= n; i++ {
			currComb = append(currComb, i)
			backtracking(i+1, remain-1, currComb)
			currComb = currComb[:len(currComb)-1]
		}
	}

	backtracking(1, k, []int{})

	return res

}
