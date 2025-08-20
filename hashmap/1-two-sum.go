func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums)) // value -> index

	for i, x := range nums {
		if j, ok := m[target-x]; ok {
			return []int{j, i}
		}
		m[x] = i
	}
	return nil
}
