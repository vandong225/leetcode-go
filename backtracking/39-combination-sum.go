package backtracking

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var performSum func(cdd []int, comb []int, sum int)

	performSum = func(cdd []int, comb []int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			dst := make([]int, len(comb))
			copy(dst, comb)
			res = append(res, dst)
			return
		}

		for i, v := range cdd {
			comb = append(comb, v)
			performSum(cdd[i:], comb, sum+v)
			comb = comb[:len(comb)-1]
		}

	}

	performSum(candidates, []int{}, 0)
	return res
}
