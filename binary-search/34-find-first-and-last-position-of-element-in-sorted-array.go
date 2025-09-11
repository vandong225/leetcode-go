package binarysearch

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var search func(findLeft bool) int

	search = func(findLeft bool) int {
		l, r, indx := 0, len(nums)-1, -1
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] == target {
				indx = mid
				if findLeft {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			if nums[mid] > target {
				r = mid - 1
			}
			if nums[mid] < target {
				l = mid + 1
			}
		}

		return indx
	}

	return []int{search(true), search(false)}
}
