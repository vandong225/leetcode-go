package backtracking

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backtracking func(src []int, permutations []int)

	backtracking = func(src []int, permutations []int) {
		if len(src) == 0 {
			dst := make([]int, len(nums))
			copy(dst, permutations)
			res = append(res, dst)
		}

		for i, v := range src {
			permutations = append(permutations, v)

			newSrc := make([]int, 0, len(src)-1)
			newSrc = append(newSrc, src[:i]...)
			newSrc = append(newSrc, src[i+1:]...)
			backtracking(newSrc, permutations)

			permutations = permutations[:len(permutations)-1]
		}
	}

	backtracking(nums, []int{})

	return res
}
