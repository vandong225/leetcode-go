package binarysearch

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (r + l) / 2
		if mid == len(nums)-1 {
			break
		}

		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
