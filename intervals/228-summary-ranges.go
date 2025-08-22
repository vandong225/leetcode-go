package intervals

import "strconv"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)

	start := 0

	for start < len(nums) {
		end := start
		for end < len(nums)-1 {
			if nums[end+1]-nums[end] != 1 {
				break
			} else {
				end++
			}
		}

		if end != start {
			res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
		} else {
			res = append(res, strconv.Itoa(nums[start]))
		}
		start = end + 1

	}

	return res

}
