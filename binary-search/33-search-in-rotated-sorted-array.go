package binarysearch

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] { // normal case
			if nums[mid] > target && nums[l] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else { // peak case
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
