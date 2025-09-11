package binarysearch

import "math"

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	min := math.MaxInt

	for l <= r {
		mid := (l + r) / 2

		if min > nums[mid] {
			min = nums[mid]
		}

		if nums[mid] >= nums[l] {
			if nums[r] < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			r = mid - 1
		}

	}

	return min
}
